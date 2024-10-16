package controller

import (
	"backend-go/internal/service"
	"backend-go/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	 systemStatusCode := uc.userService.Register("", "")
	 response.SuccessResponse(c, systemStatusCode, nil)
}

// func (uc *UserController) GetUser(c *gin.Context) {
// 	// response.SuccessResponse(c, response.Success, []string{"m10", "cr7"})
// 	response.ErrorResponse(c, response.ParamInvalid)
// }
