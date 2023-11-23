package accounts

import (
	"database/sql"
)

type TransferRepository struct {
	db *sql.DB
}

func NewTransferRepository(db *sql.DB) *TransferRepository {
	return &TransferRepository{
		db: db,
	}
}
