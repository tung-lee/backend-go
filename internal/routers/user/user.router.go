package user

import (
	"backend-go/internal/controller"
	"backend-go/internal/repo"
	"backend-go/internal/service"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// public router
	  
	// non-dependency
	urepo := repo.NewUserRepository()
	us := service.NewUserService(urepo)
	ucNonDependency := controller.NewUserController(us)
	
	userRouterPublic := router.Group("/user") 
	{
		userRouterPublic.POST("/register", ucNonDependency.Register)
		userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := router.Group("/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Auth())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/info")
	}
}