//go:build wireinject

package wire

import (
	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/repo"
	service "github.com/anle/codebase/internal/services"
	"github.com/google/wire"
)

func InitGroupDisRouterHandler() (*controller.GroupDisController, error) {
	wire.Build(
		repo.NewGroupDisRepo,
		repo.NewGroupRepo,
		repo.NewAccountRepo,
		service.NewGroupDisService,
		controller.NewGroupDisController,
	)

	return new(controller.GroupDisController), nil
}
