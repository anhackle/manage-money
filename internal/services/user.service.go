package service

import (
	"errors"

	"github.com/anle/codebase/internal/po"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/response"
	"gorm.io/gorm"
)

type IUserService interface {
	Register(userInput po.User) (int, error)
	Login(userInput po.User) (int, string, error)
}

type userService struct {
	userRepo       repo.IUserRepo
	passwordHasher IPasswordHasherService
	generateToken  IGenerateTokenService
}

// Login implements IUserService.
func (us *userService) Login(userInput po.User) (int, string, error) {
	user, err := us.userRepo.FindByEmail(userInput)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrCodeLoginFail, "", err
		}

		return response.ErrCodeInternal, "", err
	}

	err = us.passwordHasher.Compare(user.Password, userInput.Password)
	if err != nil {
		return response.ErrCodeLoginFail, "", err
	}

	accessToken, err := us.generateToken.GenerateToken(user)
	if err != nil {
		return response.ErrCodeInternal, "", err
	}

	return response.ErrCodeSuccess, accessToken, nil
}

// Register implements IUserService.
func (us *userService) Register(userInput po.User) (int, error) {
	_, err := us.userRepo.FindByEmail(userInput)
	if err == nil {
		return response.ErrCodeUserHasExists, errors.New("user existed")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		hashedPassword, err := us.passwordHasher.Hash(userInput.Password)
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

func NewUserService(userRepo repo.IUserRepo, passwordHasher IPasswordHasherService, generateToken IGenerateTokenService) IUserService {
	return &userService{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		generateToken:  generateToken,
	}
}
