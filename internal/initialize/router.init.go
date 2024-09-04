package initialize

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/middlewares"
	"github.com/anle/codebase/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	userRouter := routers.RouterGroupApp.User
	accountRouter := routers.RouterGroupApp.Account
	transactionRouter := routers.RouterGroupApp.Transaction

	MainGroup := r.Group("/v1")
	MainGroup.Use(middlewares.CORSMiddleware())

	{
		userRouter.InitUserRouter(MainGroup)
		accountRouter.InitAccountRouter(MainGroup)
		transactionRouter.InitTransactionRouter(MainGroup)
	}

	return r
}
