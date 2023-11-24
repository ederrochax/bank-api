package middlewares

import (
	"context"
	"net/http"
	"strings"

	"bank-api/internal/domain/auth/service"
	"bank-api/internal/gateways/http/utils"
)

type AuthContextKey string

var contextAccountID = AuthContextKey("account_id")

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("authorization")
		token = strings.Replace(token, "Bearer ", "", -1)

		accountID, err := service.Authorize(token)
		if err != nil {
			_ = utils.SendError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), contextAccountID, accountID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetAccountID(ctx context.Context) (string, bool) {
	tokenStr, ok := ctx.Value(contextAccountID).(string)
	return tokenStr, ok
}
