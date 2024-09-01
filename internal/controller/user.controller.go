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

	result, _ := uc.userService.Register(userInput)

	response.HandleResult(c, result, nil)
}

func (uc *UserController) Login(c *gin.Context) {
	var userInput po.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		response.ErrorResponseExternal(c, 20002, nil)
		return
	}

	result, accessToken, _ := uc.userService.Login(userInput)
	if result == response.ErrCodeSuccess {
		c.SetCookie("access-token", accessToken, 3600, "/", "localhost", true, true)
	}

	response.HandleResult(c, result, nil)
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
