package service

import (
	"github.com/anle/codebase/internal/po"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/response"
)

type IUserService interface {
	Register(userInput po.User) int
	Login(userInput po.User) int
}

type userService struct {
	userRepo       repo.IUserRepo
	passwordHasher IPasswordHasher
}

// Login implements IUserService.
func (us *userService) Login(userInput po.User) int {
	password, err := us.userRepo.Login(userInput)
	if err != nil {
		return response.ErrCodeInternal
	}

	if us.passwordHasher.Compare(password, userInput.Password) {
		return response.ErrCodeSuccess
	}

	return response.ErrCodeLoginFail
}

// Register implements IUserService.
func (us *userService) Register(userInput po.User) int {
	if ok := us.userRepo.CheckExistByEmail(userInput.Email); ok {
		return response.ErrCodeUserHasExists
	}

	hashedPassword, err := us.passwordHasher.Hash(userInput.Password)
	if err != nil {
		return response.ErrCodeInternal
	}

	userInput.Password = hashedPassword
	if ok := us.userRepo.Register(userInput); ok {
		return response.ErrCodeSuccess
	}

	return response.ErrCodeInternal
}

func NewUserService(userRepo repo.IUserRepo, passwordHasher IPasswordHasher) IUserService {
	return &userService{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
	}
}
