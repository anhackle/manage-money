//go:build wireinject

package wire

import (
	"github.com/anle/codebase/internal/controller"
	"github.com/anle/codebase/internal/repo"
	service "github.com/anle/codebase/internal/services"
	"github.com/google/wire"
)

func InitTransactionRouterHandler() (*controller.TransactionController, error) {
	wire.Build(
		repo.NewAccountRepo,
		repo.NewTransactionRepo,
		service.NewTransactionService,
		controller.NewTransactionController,
	)

	return new(controller.TransactionController), nil
}
