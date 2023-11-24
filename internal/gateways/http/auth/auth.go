package auth

import (
	"bank-api/internal/domain/auth/service"
	"bank-api/internal/gateways/http/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	authService service.Auth
	validator   *utils.StructValidator
}

func NewHandler(r *mux.Router, auth service.Auth) *Handler {
	h := &Handler{
		authService: auth,
		validator:   utils.NewValidator(),
	}

	r.HandleFunc("/login", h.Login).Methods("POST")

	return h
}
