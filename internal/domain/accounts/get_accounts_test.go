package accounts_test

import (
	"bank-api/internal/domain/entities"
	"context"
	"testing"
	"bank-api/internal/domain/accounts"

	"github.com/stretchr/testify/assert"
)

type mockGetAccountsRepository struct {
	accounts []entities.Account
	err      error
}

func (m *mockGetAccountsRepository) FindAll(ctx context.Context) ([]entities.Account, error) {
	return m.accounts, m.err
}

func Test_GetAccounts(t *testing.T) {
	mockRepo := &mockGetAccountsRepository{
		accounts: []entities.Account{
			{ID: "1", Name: "John Doe", CPF: "12345678900", Balance: 100},
			{ID: "2", Name: "Fulano De tal", CPF: "98765432100", Balance: 200},
		},
		err: nil,
	}

	uc := accounts.NewGetAccountsUC(mockRepo)
	output, err := uc.GetAccounts(context.Background())

	assert.NoError(t, err)
	assert.Len(t, output.Accounts, 2)
}
