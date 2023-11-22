package accounts

import (
	"bank-api/internal/domain/accounts"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type GetAccountBalanceUC interface {
	GetAccountBalance(ctx context.Context, accountID string) (accounts.GetAccountBalanceOutput, error)
}

func (h *AccountHandler) GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	accountID := r.URL.Query().Get("account_id")
	if accountID == "" {
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
