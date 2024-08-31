package repo

import (
	"errors"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/po"
	"gorm.io/gorm"
)

type IUserRepo interface {
	CreateUser(userInput po.User) error
	FindByEmail(userInput po.User) (po.User, error)
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
func (ur *userRepo) CreateUser(userInput po.User) error {
	result := global.Mdb.Create(&userInput)

	return result.Error
}

// Login implements IUserRepo.
func (ur *userRepo) FindByEmail(userInput po.User) (po.User, error) {
	var user po.User
	result := global.Mdb.Where("email = ?", userInput.Email).First(&user)
	if result.Error != nil {
		return po.User{}, result.Error
	}

	return user, nil
}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}
