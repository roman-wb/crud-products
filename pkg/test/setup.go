package test

import (
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

const DatabaseURL = "postgres://user:password@localhost:5432/test_db?sslmode=disable"

var db *pgxpool.Pool

func GetTables() []string {
	return []string{"products"}
}

func Setup() *pgxpool.Pool {
	var once sync.Once

	once.Do(func() {
		// Connect
		var err error
		db, err = pgxpool.Connect(context.Background(), DatabaseURL)
		if err != nil {
			log.Fatal(err)
		}

		// Truncate all tables
		Truncate()

		// TODO Close
		// defer db.Close()
	})

	return db
}

func Truncate() {
	for _, table := range GetTables() {
		_, err := db.Exec(context.Background(), `TRUNCATE TABLE `+table)
		if err != nil {
			log.Fatal(err)
		}
	}
}
