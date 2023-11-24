package service_test

import (
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"

	"bank-api/internal/domain/auth/service"
	"bank-api/internal/domain/entities"
)

func TestCreateToken(t *testing.T) {

	t.Run("Should create a valid JWT token", func(t *testing.T) {
		acc, _ := entities.NewAccount("John Duo", "12345678910", "12345678", 100)

		claims := jwt.MapClaims{}
		token, err := service.CreateToken(*acc)

		assert.NotNil(t, token)
		assert.Nil(t, err)
		if err != nil {
			return
		}

		j, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("AUTH_SECRET")), nil
		})

		assert.Nil(t, err)
		if err != nil {
			assert.True(t, j.Valid)
			assert.Equal(t, acc.ID, claims["Id"])
		}
	})
}
