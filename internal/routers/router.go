package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"

	c "backend-go/internal/controller"
	"backend-go/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before -> AA")
		c.Next()
		fmt.Println("After -> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before -> BB")
		c.Next()
		fmt.Println("After -> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before -> CC")
	c.Next()
	fmt.Println("After -> CC")
}
 
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.AuthMiddleware, AA(), BB(), CC)

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user", c.NewUserController().GetUser)
		// v1.GET("/ping/query", controller.PongWithQuery)
		// v1.GET("/ping/:name", controller.PongWithName)
		// v1.POST("/ping", controller.Pong)
		// v1.PUT("/ping", controller.Pong)
		// v1.PATCH("/ping", controller.Pong)
		// v1.DELETE("/ping", controller.Pong)
		// v1.HEAD("/ping", controller.Pong)
		// v1.OPTIONS("/ping", controller.Pong)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/ping", c.NewPongController().Pong)
	}

	return r
}
