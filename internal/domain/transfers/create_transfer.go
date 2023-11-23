package transfers

import (
	"context"
	"fmt"

	"bank-api/internal/domain/entities"
)

type createTransferRepository interface {
	Create(context.Context, *entities.Transfer) error
}

type CreateTransferUC struct {
	transfersRepo createTransferRepository
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

	err = uc.transfersRepo.Create(ctx, transfer)
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
