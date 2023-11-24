package accounts

import (
	"bank-api/internal/domain/entities"
	"context"
	"fmt"
)

type getAccountRepository interface {
	FindByID(ctx context.Context, accountID string) (entities.Account, error)
}

type getAccountUC struct {
	accountRepo getAccountRepository
}

type GetAccountOutput struct {
	Account entities.Account
}

func (uc getAccountUC) GetAccountByID(ctx context.Context, accountID string) (GetAccountOutput, error) {
	account, err := uc.accountRepo.FindByID(ctx, accountID)
	if err != nil {
		return GetAccountOutput{}, fmt.Errorf("unable to find by id: %w", err)
	}

	return GetAccountOutput{Account: account}, nil
}

func NewGetAccountUC(accountRepo getAccountRepository) getAccountUC {
	return getAccountUC{
		accountRepo: accountRepo,
	}
}
