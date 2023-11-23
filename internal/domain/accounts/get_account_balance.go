package accounts

import (
	"context"
	"fmt"
)

type getAccountBalanceRepository interface {
	FindByID(context.Context, string) (int64, error)
}

type getAccountBalanceUC struct {
	accountRepo getAccountBalanceRepository
}

type GetAccountBalanceOutput struct {
	Balance int64
}

func (uc getAccountBalanceUC) GetAccountBalance(ctx context.Context, accountID string) (GetAccountBalanceOutput, error) {
	balance, err := uc.accountRepo.FindByID(ctx, accountID)
	if err != nil {
		return GetAccountBalanceOutput{}, fmt.Errorf("unable to find by id: %w", err)
	}

	return GetAccountBalanceOutput{Balance: balance}, nil
}

func NewGetAccountBalanceUC(accountRepo getAccountBalanceRepository) getAccountBalanceUC {
	return getAccountBalanceUC{
		accountRepo: accountRepo,
	}
}
