package http

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func NewRouter(accountHandler AccountHandler, logger *zap.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/accounts", accountHandler)
	r.Get("/accounts/{accountID}/balance", accountHandler.GetAccountBalance)
	r.Post("/accounts", accountHandler.CreateAccount)

	return r
}
