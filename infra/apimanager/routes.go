package apimanager

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// NewHandler create and return new handler
func NewHandler(accountAPI AccountAPI, transactionAPI TransactionAPI) (handler *chi.Mux) {
	handler = chi.NewRouter()

	handler.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "OPTIONS", "GET"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	}).Handler)

	handler.Use(middleware.Heartbeat("/"))
	handler.Post("/accounts", accountAPI.Create)
	handler.Get("/accounts/{accountID}", accountAPI.GetByID)
	handler.Post("/transactions", transactionAPI.Create)
	handler.Post("/transactions/query", transactionAPI.Get)
	handler.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")))
	return
}
