package service

import (
	"errors"

	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/response"
	"gorm.io/gorm"
)

type IGroupDisService interface {
	ListAccountFromGroup(userID int, groupDisInput dto.GroupDisListInput) (int, []dto.GroupDisOutput, error)
	AddAccountToGroup(userID int, groupDisInput dto.GroupDisCreateInput) (int, error)
	DeleteAccountFromGroup(userID int, groupDisInput dto.GroupDisDeleteInput) (int, error)
}

type groupDisService struct {
	groupDisRepo repo.IGroupDisRepo
	groupRepo    repo.IGroupRepo
	accountRepo  repo.IAccountRepo
}

func (gds *groupDisService) ListAccountFromGroup(userID int, groupDisInput dto.GroupDisListInput) (int, []dto.GroupDisOutput, error) {
	_, err := gds.groupRepo.FindGroupByID(userID, groupDisInput.GroupID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrCodeGroupNotExist, []dto.GroupDisOutput{}, err
		}

		return response.ErrCodeInternal, []dto.GroupDisOutput{}, err
	}

	groupDistributed, err := gds.groupDisRepo.FindGroupDisByGroupID(userID, groupDisInput)
	if err != nil {
		return response.ErrCodeInternal, []dto.GroupDisOutput{}, err
	}

	return response.ErrCodeSuccess, groupDistributed, nil
}

func (gds *groupDisService) AddAccountToGroup(userID int, groupDisInput dto.GroupDisCreateInput) (int, error) {
	_, err := gds.groupRepo.FindGroupByID(userID, groupDisInput.GroupID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrCodeGroupNotExist, err
		}

		return response.ErrCodeInternal, err
	}

	_, err = gds.accountRepo.FindAccountByID(userID, groupDisInput.AccountID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrCodeAccountNotExist, err
		}

		return response.ErrCodeInternal, err
	}

	//TODO: Handle race condtion
	percentageRemain, err := gds.groupDisRepo.FindPercentageRemain(userID, groupDisInput.GroupID)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	if groupDisInput.Percentage > percentageRemain {
		return response.ErrCodePercentageExceed, nil
	}

	err = gds.groupDisRepo.Create(userID, percentageRemain, groupDisInput)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

func (gds *groupDisService) DeleteAccountFromGroup(userID int, groupDisInput dto.GroupDisDeleteInput) (int, error) {
	_, err := gds.groupRepo.FindGroupByID(userID, groupDisInput.GroupID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrCodeGroupNotExist, err
		}

		return response.ErrCodeInternal, err
	}

	_, err = gds.accountRepo.FindAccountByID(userID, groupDisInput.AccountID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrCodeAccountNotExist, err
		}

		return response.ErrCodeInternal, err
	}

	percentageRemain, err := gds.groupDisRepo.FindPercentageRemain(userID, groupDisInput.GroupID)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	err = gds.groupDisRepo.Delete(userID, percentageRemain, groupDisInput)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.ErrCodeDistributionNotExist, err
		}
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

func NewGroupDisService(groupDisRepo repo.IGroupDisRepo, groupRepo repo.IGroupRepo, accountRepo repo.IAccountRepo) IGroupDisService {
	return &groupDisService{
		groupDisRepo: groupDisRepo,
		groupRepo:    groupRepo,
		accountRepo:  accountRepo,
	}
}