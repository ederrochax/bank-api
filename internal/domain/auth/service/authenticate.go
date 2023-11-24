package service

import (
	"context"
)

func (a Auth) Authenticate(ctx context.Context, CPF, secret string) (string, error) {
	account, err := a.accountUseCase.GetAccountByCPF(ctx, CPF)
	if err != nil {
		return "", err
	}

	token, err := CreateToken(account.Account)
	if err != nil {
		return "", err
	}

	return token, nil
}
