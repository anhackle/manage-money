package repo

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/po"
)

type IUserRepo interface {
	CreateUser(userInput po.User) error
	FindByEmail(userInput po.User) (po.User, error)
	FindByUserID(userInput po.User) (po.User, error)
}

type userRepo struct{}

func (ur *userRepo) FindByUserID(userInput po.User) (po.User, error) {
	var user po.User
	result := global.Mdb.Where("id = ?", userInput.ID).First(&user)
	if result.Error != nil {
		return po.User{}, result.Error
	}

	return user, nil
}

func (ur *userRepo) CreateUser(userInput po.User) error {
	result := global.Mdb.Create(&userInput)

	return result.Error
}

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
