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
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := uc.userService.Register(userInput)

	response.HandleResult(c, result, nil)
}

func (uc *UserController) Login(c *gin.Context) {
	var userInput po.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, accessToken, _ := uc.userService.Login(userInput)

	response.HandleResult(c, result, accessToken)
}

func (uc *UserController) Profile(c *gin.Context) {
	response.HandleResult(c, response.ErrCodeSuccess, "This is profile page")
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
