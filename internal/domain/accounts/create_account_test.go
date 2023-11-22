package accounts_test

import (
	"bank-api/internal/domain/accounts"
	"bank-api/internal/domain/entities"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockCreateAccountRepository struct {
	accountID string
	err       error
}

func (m *mockCreateAccountRepository) Create(ctx context.Context, account *entities.Account) (string, error) {
	return m.accountID, m.err
}

func TestCreateAccount(t *testing.T) {
	mockRepo := &mockCreateAccountRepository{
		accountID: "123",
		err:       nil,
	}

	uc := accounts.NewCreateAccountUC(mockRepo)
	input := accounts.CreateAccountInput{
		Name:           "John Doe",
		CPF:            "12345678900",
		Secret:         "mysecret",
		InitialBalance: 100,
	}
	output, err := uc.CreateAccount(context.Background(), input)

	assert.NoError(t, err)
	assert.Equal(t, output.AccountID, "123")
}
