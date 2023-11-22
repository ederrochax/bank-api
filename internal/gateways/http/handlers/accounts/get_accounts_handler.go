package accounts

import (
	"bank-api/internal/domain/accounts"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getAccountstUC interface {
	GetAccounts(ctx context.Context) (accounts.GetAccountsOutput, error)
}

type getAccountsHandler struct {
	getAccountstUC getAccountstUC
}

func (h getAccountsHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.getAccountstUC.GetAccounts(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("error fetching accounts: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func NewGetAccountsHandler(getAccountstUC getAccountstUC) getAccountsHandler {
	return getAccountsHandler{
		getAccountstUC: getAccountstUC,
	}
}
