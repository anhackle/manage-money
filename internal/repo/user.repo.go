package repo

import (
	"errors"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/po"
	"gorm.io/gorm"
)

type IUserRepo interface {
	CheckExistByEmail(email string) bool
	Register(userInput po.User) bool
	Login(userInput po.User) (string, error)
}

type userRepo struct{}

// GetuserByEmail implements IUserRepo.
func (ur *userRepo) CheckExistByEmail(email string) bool {
	var user po.User
	result := global.Mdb.Where("email = ?", email).First(&user)
	if result.Error == nil {
		return true
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

// Register implements IUserRepo.
func (ur *userRepo) Register(userInput po.User) bool {
	result := global.Mdb.Create(&userInput)
	return result.Error == nil
}

// Login implements IUserRepo.
func (ur *userRepo) Login(userInput po.User) (string, error) {
	var user po.User
	result := global.Mdb.Where("email = ?", userInput.Email).First(&user)
	if result.Error != nil {
		return "", result.Error
	}

	return user.Password, nil
}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}
