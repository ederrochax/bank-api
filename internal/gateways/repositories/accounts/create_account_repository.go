package accounts

import (
	"bank-api/internal/domain/entities"
	"context"
)

func (r *AccountRepository) Create(ctx context.Context, account *entities.Account) (string, error) {
	query := "INSERT INTO accounts (id, name, cpf, secret, balance, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, account.ID, account.Name, account.CPF, account.Secret, account.Balance, account.CreatedAt).Scan(&account.ID)
	if err != nil {
		return "", err
	}
	return account.ID, nil
}
