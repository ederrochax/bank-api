package accounts

import (
	"bank-api/internal/domain/entities"
	"context"
)

func (r *AccountRepository) FindByCPF(ctx context.Context, cpf string) (entities.Account, error) {
	var account entities.Account
	query := "SELECT id, name, cpf, secret, balance, created_at FROM accounts WHERE cpf = $1"
	err := r.db.QueryRowContext(ctx, query, cpf).Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
	if err != nil {
		return entities.Account{}, err
	}
	return account, nil
}
