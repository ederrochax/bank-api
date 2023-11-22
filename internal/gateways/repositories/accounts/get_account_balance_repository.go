package accounts

import (
	"context"
)

func (r *AccountRepository) FindByID(ctx context.Context, accountID string) (int64, error) {
	var balance int64
	query := "SELECT balance FROM accounts WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, accountID).Scan(&balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
