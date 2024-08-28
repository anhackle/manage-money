package user

import (
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (p *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public router
	userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := Router.Group("/users")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/login", userController.Login)
	}

	//private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/get-info")
	}

}
