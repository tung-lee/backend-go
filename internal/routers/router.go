package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", Pong)
		v1.GET("/ping/query", PongWithQuery)
		v1.GET("/ping/:name", PongWithName)
		v1.POST("/ping", Pong)
		v1.PUT("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.DELETE("/ping", Pong)
		v1.HEAD("/ping", Pong)
		v1.OPTIONS("/ping", Pong)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/ping", Pong)
		v2.GET("/ping/query", PongWithQuery)
		v2.GET("/ping/:name", PongWithName)
		v2.POST("/ping", Pong)
		v2.PUT("/ping", Pong)
		v2.PATCH("/ping", Pong)
		v2.DELETE("/ping", Pong)
		v2.HEAD("/ping", Pong)
		v2.OPTIONS("/ping", Pong)
	}

	return r
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"users":   []string{"Alex", "John", "Jane"},
	})
}

func PongWithName(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong with name: " + name,
	})
}

func PongWithQuery(c *gin.Context) {
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong with query: " + uid,
	})
}
