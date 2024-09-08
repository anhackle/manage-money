package repo

import (
	"fmt"
	"time"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/po"
)

type ITokenRepo interface {
	CreateToken(user po.User, accessToken string) error
	FindUserIDByToken(accessToken string) (int, error)
}

type tokenRepo struct{}

func (gtr *tokenRepo) FindUserIDByToken(accessToken string) (int, error) {
	userID, err := global.Rdb.Get(ctx, accessToken).Int()
	if err != nil {
		return -1, err
	}

	return userID, nil
}

func (gtr *tokenRepo) CreateToken(user po.User, accessToken string) error {
	err := global.Rdb.SetEx(ctx, accessToken, user.ID, 3600*time.Second).Err()
	if err != nil {
		return err
	}

	err = global.Rdb.SetEx(ctx, fmt.Sprintf("%d", user.ID), accessToken, 3600*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func NewTokenRepo() ITokenRepo {
	return &tokenRepo{}
}
