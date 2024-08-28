package controller

import (
	"github.com/anle/codebase/internal/po"
	service "github.com/anle/codebase/internal/services"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func (uc *UserController) Register(c *gin.Context) {
	var userInput po.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		response.ErrorResponseExternal(c, 20002, nil)
		return
	}

	result := uc.userService.Register(userInput)

	switch result {
	case 20000:
		response.SuccessResponse(c, result, nil)
		return
	case 20001:
		response.ErrorResponseInternal(c, result, nil)
		return
	case 50001:
		response.ErrorResponseExternal(c, result, nil)
		return
	}
}

func (uc *UserController) Login(c *gin.Context) {
	var userInput po.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		response.ErrorResponseExternal(c, 20002, nil)
		return
	}

	result := uc.userService.Login(userInput)
	switch result {
	case 20000:
		response.SuccessResponse(c, result, nil)
		return
	case 20001:
		response.ErrorResponseInternal(c, result, nil)
		return
	case 20004:
		response.ErrorResponseExternal(c, result, nil)
		return
	}
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
