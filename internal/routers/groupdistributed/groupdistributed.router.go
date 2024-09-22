package groupdistributed

import (
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type GroupDisRouter struct{}

func (p *GroupDisRouter) InitGroupDisRouter(router *gin.RouterGroup) {
	groupDisController, _ := wire.InitGroupDisRouterHandler()
	authMiddlware, _ := wire.InitMiddlewareHandler()

	//private router
	GroupDisRouterPrivate := router.Group("/groupdistributed")
	GroupDisRouterPrivate.Use(authMiddlware.Authentication())

	{
		GroupDisRouterPrivate.GET("/", groupDisController.ListAccountFromGroup)
		GroupDisRouterPrivate.POST("/", groupDisController.AddAccountToGroup)
		GroupDisRouterPrivate.DELETE("/", groupDisController.DeleteAccountFromGroup)
	}

}
