package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/roman-wb/crud-products/internal/repos"
	"github.com/roman-wb/crud-products/internal/server"
	"go.uber.org/zap"
)

const ShutdownTimeout = 5 * time.Second

func main() {
	// Setup logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync() //nolint:errcheck

	// Load env
	err = godotenv.Load()
	if err != nil {
		logger.Sugar().Warn(err)
	}

	// Connect DB
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Sugar().Fatal("parse DB connection URL")
		return
	}
	poolConfig.ConnConfig.Logger = zapadapter.NewLogger(logger)

	db, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		logger.Sugar().Fatalf("unable to connect to database: %v", err)
	}
	defer db.Close()

	// Dependends
	repos := repos.NewRepos(db)

	// Run server
	server := server.Run(logger, repos)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	server.Shutdown(ctx)
	logger.Sugar().Infof("shutdown successful")
}
