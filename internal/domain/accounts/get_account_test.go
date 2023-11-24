package accounts_test

import (
	"bank-api/internal/domain/accounts"
	"bank-api/internal/domain/entities"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockGetAccountRepository struct {
	account entities.Account
	err     error
}

func (m *mockGetAccountRepository) FindByID(ctx context.Context, accountID string) (entities.Account, error) {
	return m.account, m.err
}

func Test_GetAccountByID(t *testing.T) {
	mockRepo := &mockGetAccountRepository{
		account: entities.Account{ID: "1", Name: "John Doe", CPF: "12345678900", Balance: 100},
		err:     nil,
	}

	uc := accounts.NewGetAccountUC(mockRepo)
	output, err := uc.GetAccountByID(context.Background(), "1")

	assert.NoError(t, err)
	assert.Equal(t, output.Account, mockRepo.account)
}
