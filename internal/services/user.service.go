package service

import (
	"errors"

	"github.com/anle/codebase/internal/po"
	"github.com/anle/codebase/internal/repo"
	hasher "github.com/anle/codebase/internal/utils/hash"
	"github.com/anle/codebase/internal/utils/token"
	"github.com/anle/codebase/response"
	"gorm.io/gorm"
)

type IUserService interface {
	Register(userInput po.User) (int, error)
	Login(userInput po.User) (int, string, error)
}

type userService struct {
	userRepo  repo.IUserRepo
	tokenRepo repo.ITokenRepo
}

func (us *userService) Login(userInput po.User) (int, string, error) {
	user, err := us.userRepo.FindByEmail(userInput)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrCodeLoginFail, "", err
		}

		return response.ErrCodeInternal, "", err
	}

	err = hasher.Compare(user.Password, userInput.Password)
	if err != nil {
		return response.ErrCodeLoginFail, "", err
	}

	accessToken, err := token.GenerateToken(user)
	if err != nil {
		return response.ErrCodeInternal, "", err
	}

	err = us.tokenRepo.CreateToken(user, accessToken)
	if err != nil {
		return response.ErrCodeInternal, "", err
	}

	return response.ErrCodeSuccess, accessToken, nil
}

func (us *userService) Register(userInput po.User) (int, error) {
	_, err := us.userRepo.FindByEmail(userInput)
	if err == nil {
		return response.ErrCodeUserHasExists, errors.New("user existed")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		hashedPassword, err := hasher.Hash(userInput.Password)
		if err != nil {
			return response.ErrCodeInternal, err
		}

		userInput.Password = hashedPassword
		err = us.userRepo.CreateUser(userInput)
		if err != nil {
			return response.ErrCodeInternal, err
		}

		return response.ErrCodeSuccess, nil
	}

	return response.ErrCodeInternal, err
}

func NewUserService(userRepo repo.IUserRepo, tokenRepo repo.ITokenRepo) IUserService {
	return &userService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}
