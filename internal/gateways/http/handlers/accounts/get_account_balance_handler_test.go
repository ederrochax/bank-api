package accounts_test

import (
	"bank-api/internal/domain/accounts"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	accounthandler "bank-api/internal/gateways/http/handlers/accounts"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGetAccountBalanceUC struct {
	mock.Mock
}

func (m *mockGetAccountBalanceUC) GetAccountBalance(ctx context.Context, accountID string) (accounts.GetAccountBalanceOutput, error) {
	args := m.Called(ctx, accountID)
	return args.Get(0).(accounts.GetAccountBalanceOutput), args.Error(1)
}

func TestGetAccountBalanceHandler_GetAccountBalance(t *testing.T) {
	mockUsecase := new(mockGetAccountBalanceUC)
	handler := accounthandler.NewGetAccountBalanceHandler(mockUsecase)
	accountID := "123"

	expectedBalance := int64(100)
	expectedOutput := accounts.GetAccountBalanceOutput{
		Balance: expectedBalance,
	}

	mockUsecase.On("GetAccountBalance", mock.Anything, accountID).Return(expectedOutput, nil)

	r := mux.NewRouter()
	r.HandleFunc("/accounts/{account_id}/balance", handler.GetAccountBalance)

	request, err := http.NewRequest("GET", fmt.Sprintf("/accounts/%s/balance", accountID), nil)
	assert.NoError(t, err)
	
	responseRecorder := httptest.NewRecorder()
	r.ServeHTTP(responseRecorder, request)


	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	var response map[string]int64
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedBalance, response["balance"])

	mockUsecase.AssertExpectations(t)
}
