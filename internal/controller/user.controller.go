package controller

import (
	"backend-go/internal/service"
	"backend-go/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	// response.SuccessResponse(c, response.Success, []string{"m10", "cr7"})
	response.ErrorResponse(c, response.ParamInvalid)
}
