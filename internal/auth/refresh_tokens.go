package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() string {

	randomToken := make([]byte, 32)

	rand.Read(randomToken)
	tokenString := hex.EncodeToString(randomToken)
	return tokenString
}
