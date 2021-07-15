package server

import (
	"net/http"
	"os"
	"time"

	defaultLogger "log"

	"github.com/roman-wb/crud-products/internal/repos"
	"go.uber.org/zap"
)

const Timeout = 5 * time.Second

func Run(logger *zap.Logger, repos *repos.Repos) *http.Server {
	router := NewRouter(logger, repos)
	server := http.Server{
		Addr:         os.Getenv("LISTEN_ADDR"),
		Handler:      router,
		WriteTimeout: Timeout,
		ReadTimeout:  Timeout,
		ErrorLog:     defaultLogger.New(&wrapperZapWriter{logger.Sugar()}, "", 0),
	}

	go func() {
		logger.Sugar().Infof("listen server on %s", server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Sugar().Fatal(err)
		}
	}()

	return &server
}

type wrapperZapWriter struct {
	logger *zap.SugaredLogger
}

func (fw *wrapperZapWriter) Write(p []byte) (n int, err error) {
	fw.logger.Errorw(string(p))
	return len(p), nil
}
