package token

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/anle/codebase/internal/po"
)

func GenerateToken(user po.User) (string, error) {
	key := make([]byte, 30)

	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	accessToken := hex.EncodeToString(key)

	return accessToken, nil
}
