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
func ValidateToken(tokenStr, secretKey string, withClaimValidation bool) (uuid.UUID, string, error) {
	var (
		key    = []byte(secretKey)
		claims = jwt.MapClaims{}
		token  *jwt.Token
		err    error
	)

	if withClaimValidation {
		token, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})
	} else {
		token, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		}, jwt.WithoutClaimsValidation())
	}
	if err != nil {
		return uuid.Nil, "", nil
	}
	if !token.Valid {
		return uuid.Nil, "", errors.New("invalid token")
	}
	userId, err := uuid.Parse(claims["id"].(string))
	username := claims["username"].(string)
	if err != nil {
		return uuid.Nil, "", errors.New("Error on the parsing")
	}
	return userId, username, nil
}
