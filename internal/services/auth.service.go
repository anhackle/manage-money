package service

import (
	"github.com/anle/codebase/internal/repo"
)

type IAuthService interface {
	Authentication(accessToken string) (int, error)
}

type authService struct {
	tokenRepo repo.ITokenRepo
}

func (as *authService) Authentication(accessToken string) (int, error) {
	userID, err := as.tokenRepo.FindUserIDByToken(accessToken)
	if err != nil {
		return -1, err
	}

	return userID, nil
}

func NewAuthService(tokenRepo repo.ITokenRepo) IAuthService {
	return &authService{
		tokenRepo: tokenRepo,
	}
}
