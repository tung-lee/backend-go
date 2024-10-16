package admin

import "github.com/gin-gonic/gin"

type AdminRouter struct {}

func (ar *AdminRouter) InitAdminRouter(router *gin.RouterGroup) {
	// public router
	adminRouterPublic := router.Group("/admin") 
	{
		adminRouterPublic.POST("/login")
	}

	// private router
	adminRouterPrivate := router.Group("/admin")
	// adminRouterPrivate.Use(Limiter())
	// adminRouterPrivate.Use(Auth())
	// adminRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/active-user")
	}
}