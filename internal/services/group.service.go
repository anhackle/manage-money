package service

import (
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/response"
)

type IGroupService interface {
	ListGroup(userID int, groupInput *dto.GroupListInput) (int, []dto.GroupOutput, error)
	CreateGroup(userID int, account dto.GroupCreateInput) (int, error)
	UpdateGroup(userID int, account dto.GroupUpdateInput) (int, error)
	DeleteGroup(userID int, account dto.GroupDeleteInput) (int, error)
}

type groupService struct {
	groupRepo repo.IGroupRepo
}

func (gs *groupService) CreateGroup(userID int, groupInput dto.GroupCreateInput) (int, error) {
	err := gs.groupRepo.CreateGroup(userID, groupInput)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

func (gs *groupService) ListGroup(userID int, groupInput *dto.GroupListInput) (int, []dto.GroupOutput, error) {
	accounts, err := gs.groupRepo.FindGroupByUserID(userID, groupInput)
	if err != nil {
		return response.ErrCodeInternal, []dto.GroupOutput{}, err
	}

	return response.ErrCodeSuccess, accounts, nil
}

func (gs *groupService) DeleteGroup(userID int, groupInput dto.GroupDeleteInput) (int, error) {
	result, err := gs.groupRepo.DeleteGroup(userID, groupInput)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	if result.RowsAffected == 0 {
		return response.ErrCodeGroupNotExist, err
	}

	return response.ErrCodeSuccess, nil
}

func (gs *groupService) UpdateGroup(userID int, groupInput dto.GroupUpdateInput) (int, error) {
	result, err := gs.groupRepo.UpdateGroup(userID, groupInput)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	if result.RowsAffected == 0 {
		return response.ErrCodeGroupNotExist, err
	}

	return response.ErrCodeSuccess, nil
}

func NewGroupService(groupRepo repo.IGroupRepo) IGroupService {
	return &groupService{
		groupRepo: groupRepo,
	}
}
