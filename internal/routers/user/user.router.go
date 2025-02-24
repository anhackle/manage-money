package user

import (
	"github.com/anle/codebase/internal/middlewares"
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (p *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	//public router
	userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := router.Group("/users")

	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/login", userController.Login)
	}

	//private router
	userRouterPrivate := router.Group("/users")
	userRouterPrivate.Use(middlewares.AuthMiddleware())

	{
		userRouterPrivate.GET("/profile", userController.Profile)
	}

}
