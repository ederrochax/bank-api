package accounts

import (
	"bank-api/internal/domain/entities"
	"context"
	"fmt"
)

type getAccountsRepository interface {
	FindAll(context.Context) ([]entities.Account, error)
}

type getAccountsUC struct {
	accountRepo getAccountsRepository
}

type GetAccountsOutput struct {
	Accounts []entities.Account
}

func (uc getAccountsUC) GetAccounts(ctx context.Context) (GetAccountsOutput, error) {
	accounts, err := uc.accountRepo.FindAll(ctx)
	if err != nil {
		return GetAccountsOutput{}, fmt.Errorf("unable to find all: %w", err)
	}

	return GetAccountsOutput{Accounts: accounts}, nil
}

func NewGetAccountsUC(accountRepo getAccountsRepository) getAccountsUC {
	return getAccountsUC{
		accountRepo: accountRepo,
	}
}
