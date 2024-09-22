package user

import (
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (p *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	//public router
	userController, _ := wire.InitUserRouterHandler()
	authMiddlware, _ := wire.InitMiddlewareHandler()

	userRouterPublic := router.Group("/users")

	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/login", userController.Login)
	}

	//private router
	userRouterPrivate := router.Group("/users")
	userRouterPrivate.Use(authMiddlware.Authentication())

	{
		userRouterPrivate.GET("/profile", userController.Profile)
	}

}
