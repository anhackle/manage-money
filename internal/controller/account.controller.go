package controller

import (
	service "github.com/anle/codebase/internal/services"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService service.IAccountService
}

func (ac *AccountController) ListAccount(c *gin.Context) {
	panic("")
}

func (ac *AccountController) CreateAccount(c *gin.Context) {
	panic("")
}
func (ac *AccountController) UpdateAccount(c *gin.Context) {
	panic("")
}
func (ac *AccountController) DeleteAccount(c *gin.Context) {
	panic("")
}

func NewAccountController(accountService service.IAccountService) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}
