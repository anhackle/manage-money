package controller

import (
	"github.com/anle/codebase/internal/dto"
	service "github.com/anle/codebase/internal/services"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type GroupDisController struct {
	groupDisService service.IGroupDisService
}

func (gdc *GroupDisController) ListAccountFromGroup(c *gin.Context) {
	var (
		groupDisListInput dto.GroupDisListinput
		userID            = c.GetInt("userID")
	)
	if err := c.ShouldBindJSON(&groupDisListInput); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, groupDistributeds, _ := gdc.groupDisService.ListAccountFromGroup(userID, groupDisListInput)

	response.HandleResult(c, result, groupDistributeds)
}

func (gdc *GroupDisController) AddAccountToGroup(c *gin.Context) {
	var (
		groupDisCreateInput dto.GroupDisCreateInput
		userID              = c.GetInt("userID")
	)

	if err := c.ShouldBindJSON(&groupDisCreateInput); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := gdc.groupDisService.AddAccountToGroup(userID, groupDisCreateInput)

	response.HandleResult(c, result, nil)
}

func (gdc *GroupDisController) DeleteAccountFromGroup(c *gin.Context) {
	var (
		groupDisDeleteInput dto.GroupDisDeleteInput
		userID              = c.GetInt("userID")
	)

	if err := c.ShouldBindJSON(&groupDisDeleteInput); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := gdc.groupDisService.DeleteAccountFromGroup(userID, groupDisDeleteInput)

	response.HandleResult(c, result, nil)
}

func NewGroupDisController(groupDisService service.IGroupDisService) *GroupDisController {
	return &GroupDisController{
		groupDisService: groupDisService,
	}
}
