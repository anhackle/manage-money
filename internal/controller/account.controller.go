package controller

import (
	"github.com/anle/codebase/internal/dto"
	service "github.com/anle/codebase/internal/services"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService service.IAccountService
}

func (ac *AccountController) ListAccount(c *gin.Context) {
	var (
		accountInput *dto.AccountListInput
		userID       = c.GetInt("userID")
		accounts     []dto.AccountOutput
	)

	result, accounts, err := ac.accountService.ListAccount(userID, accountInput)
	if err != nil {
		response.ErrorResponseInternal(c, response.ErrCodeInternal, nil)
		return
	}

	response.HandleResult(c, result, accounts)
}

func (ac *AccountController) CreateAccount(c *gin.Context) {
	var (
		accountInput dto.AccountCreateInput
		userID       = c.GetInt("userID")
	)

	if err := c.ShouldBindJSON(&accountInput); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, err := ac.accountService.CreateAccount(userID, accountInput)
	if err != nil {
		response.ErrorResponseInternal(c, response.ErrCodeInternal, nil)
		return
	}

	response.HandleResult(c, result, nil)

}
func (ac *AccountController) UpdateAccount(c *gin.Context) {
	var (
		accountInput dto.AccountUpdateInput
		userID       = c.GetInt("userID")
	)

	if err := c.ShouldBindJSON(&accountInput); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, err := ac.accountService.UpdateAccount(userID, accountInput)
	if err != nil {
		response.ErrorResponseInternal(c, response.ErrCodeInternal, nil)
		return
	}

	response.HandleResult(c, result, nil)
}
func (ac *AccountController) DeleteAccount(c *gin.Context) {
	var (
		accountInput dto.AccountDeleteInput
		userID       = c.GetInt("userID")
	)

	if err := c.ShouldBindJSON(&accountInput); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, err := ac.accountService.DeleteAccount(userID, accountInput)
	if err != nil {
		response.ErrorResponseInternal(c, response.ErrCodeInternal, nil)
		return
	}

	response.HandleResult(c, result, nil)

}

func NewAccountController(accountService service.IAccountService) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}
