package accounts

import (
	"bank-api/internal/domain/entities"
	"context"
)

func (r *AccountRepository) FindAll(ctx context.Context) ([]entities.Account, error) {
	query := "SELECT id, name, cpf, secret, balance, created_at FROM accounts"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []entities.Account
	for rows.Next() {
		var account entities.Account
		if err := rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
