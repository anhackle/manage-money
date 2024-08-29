package repo

import (
	"fmt"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/po"
)

type IGenerateTokenRepo interface {
	CreateToken(user po.User, accessToken string) error
	FindUserByToken(accessToken string) (po.User, error)
}

type generateTokenRepo struct{}

// FindToken implements IGenerateTokenRepo.
func (gtr *generateTokenRepo) FindUserByToken(accessToken string) (po.User, error) {
	var token po.Token
	result := global.Mdb.Where("token = ?", accessToken).Preload("User").First(&token)
	if result.Error != nil {
		return po.User{}, result.Error
	}
	fmt.Println(token)

	return token.User, nil
}

// CreateToken implements IGenerateTokenRepo.
func (gtr *generateTokenRepo) CreateToken(user po.User, accessToken string) error {
	var token = po.Token{
		Token:  accessToken,
		UserID: user.ID,
	}
	result := global.Mdb.Create(&token)

	return result.Error
}

func NewGenerateTokenRepo() IGenerateTokenRepo {
	return &generateTokenRepo{}
}
