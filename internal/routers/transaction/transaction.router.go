package transaction

import (
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type TransactionRouter struct{}

func (p *TransactionRouter) InitTransactionRouter(router *gin.RouterGroup) {
	transactionController, _ := wire.InitTransactionRouterHandler()
	authMiddlware, _ := wire.InitMiddlewareHandler()

	//private router
	TransactionRouterPrivate := router.Group("/transactions")
	TransactionRouterPrivate.Use(authMiddlware.Authentication())

	{
		TransactionRouterPrivate.GET("/", transactionController.ListTransaction)
		TransactionRouterPrivate.POST("/", transactionController.CreateTransaction)
	}

}
