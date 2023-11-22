package accounts_test

import (
	"bank-api/internal/domain/accounts"
	"bank-api/internal/domain/entities"
	accounthandler "bank-api/internal/gateways/http/handlers/accounts"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGetAccountsUC struct {
	mock.Mock
}

func (m *mockGetAccountsUC) GetAccounts(ctx context.Context) (accounts.GetAccountsOutput, error) {
	args := m.Called(ctx)
	return args.Get(0).(accounts.GetAccountsOutput), args.Error(1)
}

func TestGetAccountsHandler_GetAccounts(t *testing.T) {
	t.Parallel()

	mockUC := new(mockGetAccountsUC)
	handler := accounthandler.NewGetAccountsHandler(mockUC)

	req := httptest.NewRequest("GET", "/accounts", nil)
	res := httptest.NewRecorder()

	expectedOutput := accounts.GetAccountsOutput{
		Accounts: []entities.Account{
			{ID: "1", Name: "John Doe", CPF: "12345678900", Balance: 100},
			{ID: "2", Name: "Fulano De tal", CPF: "98765432100", Balance: 200},
		},
	}
	mockUC.On("GetAccounts", mock.Anything).Return(expectedOutput, nil)

	handler.GetAccounts(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var responseBody accounts.GetAccountsOutput
	err := json.Unmarshal(res.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, expectedOutput.Accounts, responseBody.Accounts)
	mockUC.AssertCalled(t, "GetAccounts", mock.Anything)
}

func TestGetAccountsHandler_GetAccounts_Error(t *testing.T) {
	mockUC := new(mockGetAccountsUC)
	handler := accounthandler.NewGetAccountsHandler(mockUC)

	req := httptest.NewRequest("GET", "/accounts", nil)
	res := httptest.NewRecorder()

	expectedError := errors.New("some error")
	mockUC.On("GetAccounts", mock.Anything).Return(accounts.GetAccountsOutput{}, expectedError)

	handler.GetAccounts(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code)
	mockUC.AssertCalled(t, "GetAccounts", mock.Anything)
}
