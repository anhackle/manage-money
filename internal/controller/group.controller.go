package controller

import (
	"github.com/anle/codebase/internal/dto"
	service "github.com/anle/codebase/internal/services"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type GroupController struct {
	groupService service.IGroupService
}

func (gc *GroupController) ListGroup(c *gin.Context) {
	var (
		groupInput *dto.GroupListInput
		userID     = c.GetInt("userID")
		groups     []dto.GroupOutput
	)

	if err := c.ShouldBindJSON(&groupInput); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, groups, err := gc.groupService.ListGroup(userID, groupInput)
	if err != nil {
		response.ErrorResponseInternal(c, response.ErrCodeInternal, nil)
		return
	}

	response.HandleResult(c, result, groups)
}

func (gc *GroupController) CreateGroup(c *gin.Context) {
	var (
		groupInput dto.GroupCreateInput
		userID     = c.GetInt("userID")
	)

	if err := c.ShouldBindJSON(&groupInput); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, err := gc.groupService.CreateGroup(userID, groupInput)
	if err != nil {
		response.ErrorResponseInternal(c, response.ErrCodeInternal, nil)
		return
	}

	response.HandleResult(c, result, nil)

}
func (gc *GroupController) UpdateGroup(c *gin.Context) {
	var (
		groupInput dto.GroupUpdateInput
		userID     = c.GetInt("userID")
	)

	if err := c.ShouldBindJSON(&groupInput); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, err := gc.groupService.UpdateGroup(userID, groupInput)
	if err != nil {
		response.ErrorResponseInternal(c, response.ErrCodeInternal, nil)
		return
	}

	response.HandleResult(c, result, nil)
}
func (gc *GroupController) DeleteGroup(c *gin.Context) {
	var (
		groupInput dto.GroupDeleteInput
		userID     = c.GetInt("userID")
	)

	if err := c.ShouldBindJSON(&groupInput); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, err := gc.groupService.DeleteGroup(userID, groupInput)
	if err != nil {
		response.ErrorResponseInternal(c, response.ErrCodeInternal, nil)
		return
	}

	response.HandleResult(c, result, nil)

}

func NewGroupController(groupService service.IGroupService) *GroupController {
	return &GroupController{
		groupService: groupService,
	}
}
