package accounts

import (
	"bank-api/internal/domain/entities"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockGetAccountBalanceRepository struct {
	account *entities.Account
	err     error
}

func (m *mockGetAccountBalanceRepository) FindByID(ctx context.Context, accountID string) (float64, error) {
	return m.account.Balance, m.err
}

func TestGetAccountBalance(t *testing.T) {
	mockRepo := &mockGetAccountBalanceRepository{
		account: &entities.Account{ID: "1", Name: "John Doe", CPF: "12345678900", Balance: 100.0},
		err:     nil,
	}

	uc := NewGetAccountBalanceUC(mockRepo)
	output, err := uc.GetAccountBalance(context.Background(), "1")

	assert.NoError(t, err)
	assert.Equal(t, output.Balance, 100.0)
}
