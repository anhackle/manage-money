package controller

import service "github.com/anle/codebase/internal/services"

type TransactionController struct {
	transactionService service.ITransactionService
}

func NewTransactionController(transactionSerivce service.ITransactionService) *TransactionController {
	return &TransactionController{
		transactionService: transactionSerivce,
	}
}
