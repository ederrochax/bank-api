package accounts_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bank-api/internal/domain/accounts"
	accounthandler "bank-api/internal/gateways/http/handlers/accounts"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCreateAccountUC struct {
	mock.Mock
}

func (m *mockCreateAccountUC) CreateAccount(ctx context.Context, input accounts.CreateAccountInput) (accounts.CreateAccountOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(accounts.CreateAccountOutput), args.Error(1)
}

func TestCreateAccountHandler(t *testing.T) {
	mockUC := new(mockCreateAccountUC)
	mockUC.On("CreateAccount", mock.Anything, mock.Anything).Return(accounts.CreateAccountOutput{"account_id"}, nil)

	handler := accounthandler.NewCreateAccountHandler(mockUC)

	requestBody := map[string]interface{}{
		"name":            "John Doe",
		"cpf":             "12345678900",
		"secret":          "mysecret",
		"initial_balance": 1000,
	}

	requestJSON, err := json.Marshal(requestBody)
	assert.NoError(t, err)

	request := httptest.NewRequest("POST", "/accounts", bytes.NewBuffer(requestJSON))
	request.Header.Set("Content-Type", "application/json")

	responseRecorder := httptest.NewRecorder()

	handler.CreateAccount(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	var responseBody map[string]string
	json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody)

	mockUC.AssertExpectations(t)
	assert.Equal(t, "account_id", responseBody["account_id"])
}
