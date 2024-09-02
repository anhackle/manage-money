package controller

import (
	"github.com/anle/codebase/internal/dto"
	service "github.com/anle/codebase/internal/services"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService service.IAccountService
}

func (ac *AccountController) ListAccount() (c *gin.Context) {
	panic("")
}

func (ac *AccountController) CreateAccount() (c *gin.Context) {
	panic("")
}
func (ac *AccountController) UpdateAccount(account dto.Account) (c *gin.Context) {
	panic("")
}
func (ac *AccountController) DeleteAccount(account dto.Account) (c *gin.Context) {
	panic("")
}

func NewAccountController(accountService service.IAccountService) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}
