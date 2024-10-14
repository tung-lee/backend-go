package middlewares

import (
	"backend-go/pkg/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	fmt.Println("Before -> Auth Middleware")

	token := c.GetHeader("Authorization")

	if token != "valid-token" {
		response.ErrorResponse(c, response.TokenInvalid)
		c.Abort()
		return
	}

	c.Next()
	fmt.Println("After -> Auth Middleware")

}
