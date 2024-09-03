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

// Injectors from account.wire.go:

func InitAccountRouterHandler() (*controller.AccountController, error) {
	iAccountRepo := repo.NewAccountRepo()
	iAccountService := service.NewAccountService(iAccountRepo)
	accountController := controller.NewAccountController(iAccountService)
	return accountController, nil
}

// Injectors from transaction.wire.go:

func InitTransactionRouterHandler() (*controller.TransactionController, error) {
	iAccountRepo := repo.NewAccountRepo()
	iTransactionRepo := repo.NewTransactionRepo(iAccountRepo)
	iTransactionService := service.NewTransactionService(iTransactionRepo, iAccountRepo)
	transactionController := controller.NewTransactionController(iTransactionService)
	return transactionController, nil
}

// Injectors from user.wire.go:

func InitUserRouterHandler() (*controller.UserController, error) {
	iUserRepo := repo.NewUserRepo()
	iTokenRepo := repo.NewTokenRepo()
	iUserService := service.NewUserService(iUserRepo, iTokenRepo)
	userController := controller.NewUserController(iUserService)
	return userController, nil
}
