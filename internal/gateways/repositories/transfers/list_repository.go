package accounts

import (
	"bank-api/internal/domain/entities"
	"context"
)

func (r TransferRepository) ListTransfersByAccountID(ctx context.Context, ID string) ([]entities.Transfer, error) {
	statement := `
		SELECT 
			id,
			account_origin_id,
			account_destination_id,
			amount,
			created_at
		FROM 
			transfer
		WHERE account_origin_id=$1 OR account_destination_id=$1`

	rows, err := r.db.QueryContext(ctx, statement, ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transfers := []entities.Transfer{}

	for rows.Next() {
		var t entities.Transfer
		if err := rows.Scan(
			&t.ID,
			&t.AccountOriginID,
			&t.AccountDestinationID,
			&t.Amount,
			&t.CreatedAt,
		); err != nil {
			return nil, err
		}

		transfers = append(transfers, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transfers, nil
}
