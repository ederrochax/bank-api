package auth

import (
	"bank-api/internal/gateways/http/utils"
	"net/http"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqBody LoginRequest
	err := utils.Decode(r, &reqBody)
	if err != nil {
		_ = utils.SendError(w, utils.ErrDecode, http.StatusBadRequest)
		return
	}

	var validationErr ValidationErrorResponse
	err = h.validator.Validate(reqBody, &validationErr)
	if err != nil {
		_ = utils.Send(w, validationErr, http.StatusBadRequest)
		return
	}

	token, err := h.authService.Authenticate(r.Context(), reqBody.CPF, reqBody.Secret)
	if err != nil {
		_ = utils.SendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	_ = utils.Send(
		w,
		LoginResponse{Token: token},
		http.StatusOK,
	)
}
