package controller

import (
	service "github.com/anle/codebase/internal/services"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionService service.ITransactionService
}

func (tc *TransactionController) ListTransaction(c *gin.Context) {
	panic("")
}

func (tc *TransactionController) CreateTransaction(c *gin.Context) {
	panic("")
}

func NewTransactionController(transactionSerivce service.ITransactionService) *TransactionController {
	return &TransactionController{
		transactionService: transactionSerivce,
	}
}
