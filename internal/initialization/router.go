package initialization

import (
	"backend-go/global"
	"backend-go/internal/routers"

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

	// middlewares
	// r.Use() // logging
	// r.Use() // cors
	// r.Use() // limiter global

	adminRouter := routers.RouterGroupApp.Admin
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/health") // health check
	}
	{
		userRouter.InitProductRouter(MainGroup)
		userRouter.InitUserRouter(MainGroup)
	}
	{
		adminRouter.InitAdminRouter(MainGroup)
		adminRouter.InitUserRouter(MainGroup)
	}
	
	return r
}
