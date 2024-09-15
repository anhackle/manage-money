package controller

import (
	"github.com/anle/codebase/internal/dto"
	service "github.com/anle/codebase/internal/services"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionService service.ITransactionService
}

func (tc *TransactionController) ListTransaction(c *gin.Context) {
	var (
		userID       = c.GetInt("userID")
		transactions []dto.TransOutput
	)
	result, transactions, err := tc.transactionService.ListTransaction(userID)
	if err != nil {
		response.ErrorResponseInternal(c, response.ErrCodeInternal, nil)
		return
	}

	response.HandleResult(c, result, transactions)
}

func (tc *TransactionController) CreateTransaction(c *gin.Context) {
	var (
		transactionInput dto.TransCreateInput
		userID           = c.GetInt("userID")
	)

	if err := c.ShouldBindJSON(&transactionInput); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := tc.transactionService.MakeTransaction(userID, transactionInput)

	response.HandleResult(c, result, nil)

}

func NewTransactionController(transactionSerivce service.ITransactionService) *TransactionController {
	return &TransactionController{
		transactionService: transactionSerivce,
	}
}
