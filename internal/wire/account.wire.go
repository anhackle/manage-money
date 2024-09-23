//go:build wireinject

package wire

import (
	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/repo"
	service "github.com/anle/codebase/internal/services"
	"github.com/google/wire"
)

func InitAccountRouterHandler() (*controller.AccountController, error) {
	wire.Build(
		repo.NewAccountRepo,
		service.NewAccountService,
		controller.NewAccountController,
	)

	return new(controller.AccountController), nil
}
