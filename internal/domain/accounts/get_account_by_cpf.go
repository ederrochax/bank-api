package accounts

import (
	"bank-api/internal/domain/entities"
	"context"
	"fmt"
)

type getAccountByCPFRepository interface {
	FindByCPF(ctx context.Context, cpf string) (entities.Account, error)
}

type getAccountByCPFUC struct {
	accountRepo getAccountByCPFRepository
}

type GetAccountByCPFOutput struct {
	Account entities.Account
}

func (uc getAccountByCPFUC) GetAccountByCPF(ctx context.Context, cpf string) (GetAccountByCPFOutput, error) {
	account, err := uc.accountRepo.FindByCPF(ctx, cpf)
	if err != nil {
		return GetAccountByCPFOutput{}, fmt.Errorf("unable to find by id: %w", err)
	}

	return GetAccountByCPFOutput{Account: account}, nil
}

func NewGetAccountByCPFUC(accountRepo getAccountByCPFRepository) getAccountByCPFUC {
	return getAccountByCPFUC{
		accountRepo: accountRepo,
	}
}
