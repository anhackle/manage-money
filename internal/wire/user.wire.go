//go:build wireinject

package wire

import (
	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/repo"
	service "github.com/anle/codebase/internal/services"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepo,
		repo.NewTokenRepo,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
