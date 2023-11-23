package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidAmount       = fmt.Errorf("the amount must be greater than 0")
	ErrOrigAccEqualDestAcc = fmt.Errorf("the destination account can't be equal the origin account")
)

type Transfer struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               int64
	CreatedAt            time.Time
}

func NewTransfer(accountOriginID, accountDestinationID string, amount int64) (*Transfer, error) {
	transfer := &Transfer{
		ID:                   uuid.NewString(),
		AccountOriginID:      accountOriginID,
		AccountDestinationID: accountDestinationID,
		Amount:               amount,
		CreatedAt:            time.Now(),
	}

	err := transfer.Validate()
	if err != nil {
		return nil, err
	}

	return transfer, nil
}

func (e Transfer) Validate() error {
	if e.Amount <= 0 {
		return ErrInvalidAmount
	}
	if e.AccountOriginID == e.AccountDestinationID {
		return ErrOrigAccEqualDestAcc
	}

	return nil
}
