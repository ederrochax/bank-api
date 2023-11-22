package accounts_test

import (
	"bank-api/internal/domain/accounts"
	"bank-api/internal/domain/entities"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockGetAccountBalanceRepository struct {
	account *entities.Account
	err     error
}

func (m *mockGetAccountBalanceRepository) FindByID(ctx context.Context, accountID string) (int64, error) {
	return m.account.Balance, m.err
}

func TestGetAccountBalance(t *testing.T) {
	mockRepo := &mockGetAccountBalanceRepository{
		account: &entities.Account{ID: "1", Name: "John Doe", CPF: "12345678900", Balance: 100},
		err:     nil,
	}

	uc := accounts.NewGetAccountBalanceUC(mockRepo)
	output, err := uc.GetAccountBalance(context.Background(), "1")

	assert.NoError(t, err)
	assert.Equal(t, output.Balance, 100)
}
