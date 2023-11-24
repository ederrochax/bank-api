package accounts_test

import (
	"bank-api/internal/domain/accounts"
	"bank-api/internal/domain/entities"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockGetAccountByCPFRepository struct {
	account entities.Account
	err     error
}

func (m *mockGetAccountByCPFRepository) FindByCPF(ctx context.Context, accountID string) (entities.Account, error) {
	return m.account, m.err
}

func Test_GetAccountByCPF(t *testing.T) {
	mockRepo := &mockGetAccountByCPFRepository{
		account: entities.Account{ID: "1", Name: "John Doe", CPF: "12345678900", Balance: 100},
		err:     nil,
	}

	uc := accounts.NewGetAccountByCPFUC(mockRepo)
	output, err := uc.GetAccountByCPF(context.Background(), "1")

	assert.NoError(t, err)
	assert.Equal(t, output.Account, mockRepo.account)
}
