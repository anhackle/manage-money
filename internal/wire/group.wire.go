//go:build wireinject

package wire

import (
	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/repo"
	service "github.com/anle/codebase/internal/services"
	"github.com/google/wire"
)

func InitGroupRouterHandler() (*controller.GroupController, error) {
	wire.Build(
		repo.NewGroupRepo,
		service.NewGroupService,
		controller.NewGroupController,
	)

	return new(controller.GroupController), nil
}
