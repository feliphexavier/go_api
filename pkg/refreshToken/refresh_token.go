package refreshtoken

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRefreshToken() (string, error) {
	str := make([]byte, 18)
	_, err := rand.Read(str)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(str), nil
}
