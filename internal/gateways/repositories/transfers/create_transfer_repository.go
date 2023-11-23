package accounts

import (
	"bank-api/internal/domain/entities"
	"context"
)

func (r *TransferRepository) Create(ctx context.Context, transfer *entities.Transfer) error {
	query := `INSERT INTO transfers (id, account_origin_id, account_destination_id, amount, created_at)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id`
	err := r.db.QueryRowContext(ctx, query, transfer.ID, transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount, transfer.CreatedAt).Scan(&transfer.ID)
	if err != nil {
		return err
	}
	return nil
}
