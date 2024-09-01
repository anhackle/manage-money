package repo

import (
	"fmt"
	"time"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/po"
)

type ITokenRepo interface {
	CreateToken(user po.User, accessToken string) error
	FindUserByToken(accessToken string) (po.User, error)
}

type tokenRepo struct{}

func (gtr *tokenRepo) FindUserByToken(accessToken string) (po.User, error) {
	var token po.Token
	result := global.Mdb.Where("token = ?", accessToken).Preload("User").First(&token)
	if result.Error != nil {
		return po.User{}, result.Error
	}
	fmt.Println(token)

	return token.User, nil
}

func (gtr *tokenRepo) CreateToken(user po.User, accessToken string) error {
	// TODO: Add this token to Redis instead of Mysql
	err := global.Rdb.SetEx(ctx, accessToken, user.Email, 3600*time.Second).Err()
	if err != nil {
		return err
	}

	err = global.Rdb.SetEx(ctx, user.Email, accessToken, 3600*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func NewTokenRepo() ITokenRepo {
	return &tokenRepo{}
}
