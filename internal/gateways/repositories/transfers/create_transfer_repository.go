package accounts

import (
	"bank-api/internal/domain/entities"
	"context"
	"database/sql"
	"log"
)

func (r *TransferRepository) Create(ctx context.Context, input entities.TransactionInput) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			rollErr := tx.Rollback()
			if rollErr != nil {
				log.Println(err)
			}
		}
	}()

	err = r.updateAccountBalance(ctx, tx, *input.OriginAcount)
	if err != nil {
		return err
	}

	err = r.updateAccountBalance(ctx, tx, *input.DestinationAcount)
	if err != nil {
		return err
	}

	err = r.saveTransfer(ctx, tx, input.Transfer)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *TransferRepository) saveTransfer(ctx context.Context, tx *sql.Tx, transfer *entities.Transfer) error {
	query := `INSERT INTO transfers (id, account_origin_id, account_destination_id, amount, created_at)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id`
	err := tx.QueryRowContext(ctx, query, transfer.ID, transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount, transfer.CreatedAt).Scan(&transfer.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r TransferRepository) updateAccountBalance(ctx context.Context, tx *sql.Tx, acc entities.Account) error {
	statement := `
		UPDATE
			account
		SET
			balance=$1
		WHERE 
			id=$2`

	_, err := tx.ExecContext(ctx, statement, acc.Balance, acc.ID)

	return err
}
