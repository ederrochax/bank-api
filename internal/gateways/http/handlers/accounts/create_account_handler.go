package accounts

import (
	"bank-api/internal/domain/accounts"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type createAccountUC interface {
	CreateAccount(ctx context.Context, input accounts.CreateAccountInput) (accounts.CreateAccountOutput, error)
}

type createAccountHandler struct {
	createAccountUC createAccountUC
}

type CreateBodyRequest struct {
	Name           string `json:"name"`
	CPF            string `json:"cpf"`
	Secret         string `json:"secret"`
	InitialBalance int64  `json:"initial_balance"`
}

func (h createAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateBodyRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "error decoding account data", http.StatusBadRequest)
		return
	}

	input := accounts.CreateAccountInput{
		Name:           reqBody.Name,
		CPF:            reqBody.CPF,
		Secret:         reqBody.Secret,
		InitialBalance: reqBody.InitialBalance,
	}

	accountOutput, err := h.createAccountUC.CreateAccount(r.Context(), input)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating account: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"account_id": accountOutput.AccountID})
}

func NewCreateAccountHandler(createAccountUC createAccountUC) createAccountHandler {
	return createAccountHandler{
		createAccountUC: createAccountUC,
	}
}
