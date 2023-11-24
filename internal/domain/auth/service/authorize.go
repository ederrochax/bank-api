package service

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func Authorize(token string) (interface{}, error) {
	claims := jwt.MapClaims{}

	if token == "" {
		return nil, fmt.Errorf("token not provided")
	}

	j, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("AUTH_SECRET")), nil
	})

	if err != nil || !j.Valid {
		return nil, err
	}

	if id, ok := claims["Id"]; ok {
		return id, nil
	}

	return nil, fmt.Errorf("invalid token")
}
