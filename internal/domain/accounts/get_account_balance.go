package accounts

import (
	"context"
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
		return GetAccountBalanceOutput{}, err
	}

	return GetAccountBalanceOutput{Balance: balance}, nil
}

func NewGetAccountBalanceUC(accountRepo getAccountBalanceRepository) getAccountBalanceUC {
	return getAccountBalanceUC{
		accountRepo: accountRepo,
	}
}
