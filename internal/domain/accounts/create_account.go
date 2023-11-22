package accounts

import (
	"bank-api/internal/domain/entities"
	"context"
	"fmt"
	"time"
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
	if err := input.Validate(); err != nil {
		return CreateAccountOutput{}, err
	}

	newAccount := &entities.Account{
		Name:      input.Name,
		CPF:       input.CPF,
		Secret:    input.Secret,
		Balance:   input.InitialBalance,
		CreatedAt: time.Now(),
	}

	accountID, err := uc.accountRepo.Create(ctx, newAccount)
	if err != nil {
		return CreateAccountOutput{}, err
	}

	return CreateAccountOutput{AccountID: accountID}, nil
}

func (i CreateAccountInput) Validate() error {
	if i.Name == "" || i.CPF == "" || i.Secret == "" {
		return fmt.Errorf("%w: name, CPF, and secret are required", fmt.Errorf("aa"))//ErrMalformedParameters)
	}
	return nil
}

func NewCreateAccountUC(accountRepo createAccountRepository) createAccountUC {
	return createAccountUC{
		accountRepo: accountRepo,
	}
}
