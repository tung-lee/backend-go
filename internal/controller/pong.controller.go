package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct {}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	fmt.Println("-> Pong Handler")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"users":   []string{"Alex", "John", "Jane"},
	})
}

func (pc *PongController) PongWithName(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong with name: " + name,
	})
}

func (pc *PongController) PongWithQuery(c *gin.Context) {
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong with query: " + uid,
	})
}