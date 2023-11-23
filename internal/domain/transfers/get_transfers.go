package transfers

import (
	"context"
	"fmt"

	"bank-api/internal/domain/entities"
)

type getTransfersRepository interface {
	ListByAccountID(ctx context.Context, accountOriginID string) ([]entities.Transfer, error)
}

type getTransfersByAccountIDUC struct {
	transfersRepo getTransfersRepository
}

type GetTransfersInput struct {
	AccountOriginID      string
	AccountDestinationID string
	Amount               int64
}

type GetTransfersOutput struct {
	Transfers []entities.Transfer
}

func (uc getTransfersByAccountIDUC) GetTransfersByAccountID(ctx context.Context, accountOriginID string) (GetTransfersOutput, error) {
	transfers, err := uc.transfersRepo.ListByAccountID(ctx, accountOriginID)
	if err != nil {
		return GetTransfersOutput{}, fmt.Errorf("unable to list by account id: %w", err)
	}

	return GetTransfersOutput{
		Transfers: transfers,
	}, nil
}

func NewGetTransfersByAccountIDUC(transfersRepo getTransfersRepository) getTransfersByAccountIDUC {
	return getTransfersByAccountIDUC{
		transfersRepo: transfersRepo,
	}
}
