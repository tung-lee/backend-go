package admin

import "github.com/gin-gonic/gin"

type UserRouter struct {}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// public router
	userRouterPublic := router.Group("/admin/user") 
	{
		userRouterPublic.POST("/login")
	}

	// private router
	userRouterPrivate := router.Group("/admin/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Auth())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/active-user")
	}
}