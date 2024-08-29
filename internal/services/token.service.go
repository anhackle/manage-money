package service

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/anle/codebase/internal/po"
	"github.com/anle/codebase/internal/repo"
)

type IGenerateTokenService interface {
	GenerateToken(user po.User) (string, error)
}

type generateTokenService struct {
	generateTokenRepo repo.IGenerateTokenRepo
}

// GenerateToken implements IGenerateTokenService.
func (gts *generateTokenService) GenerateToken(user po.User) (string, error) {
	key := make([]byte, 30)

	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	accessToken := hex.EncodeToString(key)
	err = gts.generateTokenRepo.CreateToken(user, accessToken)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func NewGenerateTokenService(generateTokenRepo repo.IGenerateTokenRepo) IGenerateTokenService {
	return &generateTokenService{
		generateTokenRepo: generateTokenRepo,
	}
}
