//go:build wireinject

package wire

import (
	"github.com/anle/codebase/internal/middlewares"
	"github.com/anle/codebase/internal/repo"
	service "github.com/anle/codebase/internal/services"
	"github.com/google/wire"
)

func InitMiddlewareHandler() (*middlewares.AuthMiddleware, error) {
	wire.Build(
		repo.NewTokenRepo,
		service.NewAuthService,
		middlewares.NewAuthMiddleware,
	)

	return new(middlewares.AuthMiddleware), nil
}
