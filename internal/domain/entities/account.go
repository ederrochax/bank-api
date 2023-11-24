package entities

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrMalformedParameters = fmt.Errorf("malformed parameters")
)

type Account struct {
	ID        string
	Name      string
	CPF       string
	Secret    string
	Balance   int64
	CreatedAt time.Time
}

func NewAccount(name, CPF, secret string, initialBalance int64) (*Account, error) {
	account := Account{
		ID:      uuid.NewString(),
		Name:    name,
		CPF:     CPF,
		Secret:  secret,
		Balance: initialBalance,
	}
	err := account.Validate()
	hash := account.hash(account.Secret)

	account.Secret = hash

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (e Account) Validate() error {
	if e.Name == "" || e.CPF == "" || e.Secret == "" {
		return fmt.Errorf("%w: name, CPF, and secret are required", ErrMalformedParameters)
	}
	return nil
}

func (e Account) CheckWalletFunds(amount int64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if (e.Balance - amount) < 0 {
		return ErrInsufficientFunds
	}

	return nil
}

func (e *Account) DepositMoney(amount int64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	e.Balance += amount
	return nil
}

func (e *Account) WithdrawMoney(amount int64) error {
	err := e.CheckWalletFunds(amount)
	if err != nil {
		return err
	}

	e.Balance -= amount
	return nil
}

func (e Account) hash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}
