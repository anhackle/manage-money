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

	r.Use(middlewares.CORSMiddleware())

	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1")
	{
		userRouter.InitUserRouter(MainGroup)
	}

	return r
}
