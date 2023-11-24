package transfers

import (
	"context"
	"fmt"

	"bank-api/internal/domain/accounts"
	"bank-api/internal/domain/entities"
)

type createTransferRepository interface {
	Create(context.Context, entities.TransactionInput) error
}

type getAccountByIDUC interface {
	GetAccountByID(ctx context.Context, accountID string) (accounts.GetAccountOutput, error)
}

type CreateTransferUC struct {
	transfersRepo  createTransferRepository
	accountUseCase getAccountByIDUC
}

type CreateTransferInput struct {
	AccountOriginID      string
	AccountDestinationID string
	Amount               int64
}

type CreateTransferOutput struct {
	TransferID string
}

func (uc CreateTransferUC) CreateTransfer(ctx context.Context, input CreateTransferInput) (CreateTransferOutput, error) {
	transfer, err := entities.NewTransfer(input.AccountOriginID, input.AccountDestinationID, input.Amount)
	if err != nil {
		return CreateTransferOutput{}, fmt.Errorf("unable to new transfer: %w", err)
	}

	origAcc, err := uc.accountUseCase.GetAccountByID(ctx, input.AccountOriginID)
	if err != nil {
		return CreateTransferOutput{}, err
	}

	destAcc, err := uc.accountUseCase.GetAccountByID(ctx, input.AccountDestinationID)
	if err != nil {
		return CreateTransferOutput{}, err
	}

	err = origAcc.Account.WithdrawMoney(input.Amount)
	if err != nil {
		return CreateTransferOutput{}, err
	}

	err = destAcc.Account.DepositMoney(input.Amount)
	if err != nil {
		return CreateTransferOutput{}, err
	}

	err = uc.transfersRepo.Create(ctx, entities.TransactionInput{
		OriginAcount:      &origAcc.Account,
		DestinationAcount: &destAcc.Account,
		Transfer:          transfer,
	})
	if err != nil {
		return CreateTransferOutput{}, fmt.Errorf("unable to save transfer: %w", err)
	}

	return CreateTransferOutput{
		TransferID: transfer.ID,
	}, nil
}

func NewCreateTransferUC(transfersRepo createTransferRepository) CreateTransferUC {
	return CreateTransferUC{
		transfersRepo: transfersRepo,
	}
}
