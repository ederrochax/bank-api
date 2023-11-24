package accounts

import (
	"bank-api/internal/domain/accounts"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type getAccountBalanceUC interface {
	GetAccountByID(ctx context.Context, accountID string) (accounts.GetAccountOutput, error)
}

type getAccountBalanceHandler struct {
	getAccountBalanceUC getAccountBalanceUC
}

func (h getAccountBalanceHandler) GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID, ok := vars["account_id"]
	if !ok {
		http.Error(w, "parameter 'account_id' missing", http.StatusBadRequest)
		return
	}

	account, err := h.getAccountBalanceUC.GetAccountByID(r.Context(), accountID)
	if err != nil {
		http.Error(w, fmt.Sprintf("error fetching account balance: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"balance": account.Account.Balance})
}

func NewGetAccountBalanceHandler(getAccountBalanceUC getAccountBalanceUC) getAccountBalanceHandler {
	return getAccountBalanceHandler{
		getAccountBalanceUC: getAccountBalanceUC,
	}
}
