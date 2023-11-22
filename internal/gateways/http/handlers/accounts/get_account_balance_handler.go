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
	GetAccountBalance(ctx context.Context, accountID string) (accounts.GetAccountBalanceOutput, error)
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

	balance, err := h.getAccountBalanceUC.GetAccountBalance(r.Context(), accountID)
	if err != nil {
		http.Error(w, fmt.Sprintf("error fetching account balance: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"balance": balance.Balance})
}

func NewGetAccountBalanceHandler(getAccountBalanceUC getAccountBalanceUC) getAccountBalanceHandler {
	return getAccountBalanceHandler{
		getAccountBalanceUC: getAccountBalanceUC,
	}
}
