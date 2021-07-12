package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/purini-to/zapmw"
	"github.com/roman-wb/crud-products/internal/repos"
	h "github.com/roman-wb/crud-products/internal/server/handlers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewRouter(logger *zap.Logger, repos *repos.Repos) *mux.Router {
	productHandler := h.NewProductHandler(logger, repos.Product)

	router := mux.NewRouter()
	router.HandleFunc("/", h.HomeHandler).Methods("GET")
	router.HandleFunc("/health", h.HealthHandler).Methods("GET")
	router.HandleFunc("/products", productHandler.IndexHandler).Methods("GET")
	router.HandleFunc("/products", productHandler.CreateHandler).Methods("POST", "PUT", "PATCH")
	router.HandleFunc("/products/{id}", productHandler.ShowHandler).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.UpdateHandler).Methods("POST", "PUT", "PATCH")
	router.HandleFunc("/products/{id}", productHandler.DestroyHandler).Methods("DELETE")

	router.Use(handlers.RecoveryHandler())
	router.Use(
		zapmw.WithZap(logger),
		zapmw.Request(zapcore.InfoLevel, "request"),
		zapmw.Recoverer(zapcore.ErrorLevel, "recover", zapmw.RecovererDefault),
	)

	return router
}
