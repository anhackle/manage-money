package account

import (
	"github.com/anle/codebase/internal/middlewares"
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct{}

func (p *AccountRouter) InitAccountRouter(router *gin.RouterGroup) {
	accountController, _ := wire.InitAccountRouterHandler()

	//private router
	AccountRouterPrivate := router.Group("/accounts")
	AccountRouterPrivate.Use(middlewares.AuthMiddleware())

	{
		AccountRouterPrivate.GET("/", accountController.ListAccount)
		AccountRouterPrivate.POST("/", accountController.CreateAccount)
		AccountRouterPrivate.PUT("/:id", accountController.UpdateAccount)
		AccountRouterPrivate.DELETE("/:id", accountController.DeleteAccount)
	}

}
