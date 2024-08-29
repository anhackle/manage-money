// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/internal/services"
)

// Injectors from user.wire.go:

func InitUserRouterHandler() (*controller.UserController, error) {
	iUserRepo := repo.NewUserRepo()
	iPasswordHasherService := service.NewPasswordHasher()
	iGenerateTokenRepo := repo.NewGenerateTokenRepo()
	iGenerateTokenService := service.NewGenerateTokenService(iGenerateTokenRepo)
	iUserService := service.NewUserService(iUserRepo, iPasswordHasherService, iGenerateTokenService)
	userController := controller.NewUserController(iUserService)
	return userController, nil
}
