package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateToken(id uuid.UUID, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       id,
			"username": username,
			"exp":      time.Now().Add(60 * time.Minute).Unix(),
		},
	)
	key := []byte(secretKey)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", errors.New("failed to create JWT")
	}
	return tokenStr, nil
}
