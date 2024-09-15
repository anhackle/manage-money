package group

import (
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type GroupRouter struct{}

func (p *GroupRouter) InitGroupRouter(router *gin.RouterGroup) {
	groupController, _ := wire.InitGroupRouterHandler()
	authMiddlware, _ := wire.InitMiddlewareHandler()

	//private router
	GroupRouterPrivate := router.Group("/groups")
	GroupRouterPrivate.Use(authMiddlware.Authentication())

	{
		GroupRouterPrivate.GET("/", groupController.ListGroup)
		GroupRouterPrivate.POST("/", groupController.CreateGroup)
		GroupRouterPrivate.PUT("/", groupController.UpdateGroup)
		GroupRouterPrivate.DELETE("/", groupController.DeleteGroup)
	}

}
