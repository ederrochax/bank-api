package service

import (
	"bank-api/internal/domain/accounts"
	"context"
)

type Service interface {
	Authenticate(ctx context.Context, CPF, secret string) (string, error)
}

type getAccountByCPFUC interface {
	GetAccountByCPF(ctx context.Context, cpf string) (accounts.GetAccountByCPFOutput, error)
}

type Auth struct {
	accountUseCase getAccountByCPFUC
}

func NewAuthService(accUseCase getAccountByCPFUC) Service {
	return &Auth{
		accountUseCase: accUseCase,
	}
}
