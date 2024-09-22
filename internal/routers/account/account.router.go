package account

import (
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct{}

func (p *AccountRouter) InitAccountRouter(router *gin.RouterGroup) {
	accountController, _ := wire.InitAccountRouterHandler()
	authMiddlware, _ := wire.InitMiddlewareHandler()

	//private router
	AccountRouterPrivate := router.Group("/accounts")
	AccountRouterPrivate.Use(authMiddlware.Authentication())

	{
		AccountRouterPrivate.GET("/", accountController.ListAccount)
		AccountRouterPrivate.POST("/", accountController.CreateAccount)
		AccountRouterPrivate.PUT("/", accountController.UpdateAccount)
		AccountRouterPrivate.DELETE("/", accountController.DeleteAccount)
	}

}
