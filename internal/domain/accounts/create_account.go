package accounts

import (
	"bank-api/internal/domain/entities"
	"context"
)

type createAccountRepository interface {
	Create(context.Context, *entities.Account) (string, error)
}

type createAccountUC struct {
	accountRepo createAccountRepository
}

type CreateAccountInput struct {
	Name           string
	CPF            string
	Secret         string
	InitialBalance int64
}

type CreateAccountOutput struct {
	AccountID string
}

func (uc createAccountUC) CreateAccount(ctx context.Context, input CreateAccountInput) (CreateAccountOutput, error) {
	newAccount, err := entities.NewAccount(input.Name, input.CPF, input.Secret, input.InitialBalance)
	if err != nil {
		return CreateAccountOutput{}, err
	}

	accountID, err := uc.accountRepo.Create(ctx, newAccount)
	if err != nil {
		return CreateAccountOutput{}, err
	}

	return CreateAccountOutput{AccountID: accountID}, nil
}

func NewCreateAccountUC(accountRepo createAccountRepository) createAccountUC {
	return createAccountUC{
		accountRepo: accountRepo,
	}
}
