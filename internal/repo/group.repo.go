package repo

import (
	"fmt"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/po"
)

type IGroupRepo interface {
	CreateGroup(userID int, groupInput dto.GroupCreateInput) error
	FindGroupByID(userID int, groupID int) (dto.GroupOutput, error)
	FindGroupByUserID(userID int) ([]dto.GroupOutput, error)
	UpdateGroup(userID int, groupInput dto.GroupUpdateInput) error
	DeleteGroup(userID int, groupInput dto.GroupDeleteInput) error
}

type GroupRepo struct{}

// DeleteGroup implements IGroupRepo.
func (gr *GroupRepo) DeleteGroup(userID int, groupInput dto.GroupDeleteInput) error {
	result := global.Mdb.Model(&po.Account{}).Where("userID = ? AND id = ?", userID, groupInput.ID).Delete(&po.Account{})
	return result.Error
}

// FindGroupByID implements IGroupRepo.
func (gr *GroupRepo) FindGroupByID(userID int, groupID int) (dto.GroupOutput, error) {
	var group dto.GroupOutput
	result := global.Mdb.Model(&po.Account{}).Where("userID = ? AND id = ?", userID, groupID).First(&group)
	if result.Error != nil {
		return dto.GroupOutput{}, result.Error
	}

	return group, nil
}

// FindGroupByUserID implements IGroupRepo.
func (gr *GroupRepo) FindGroupByUserID(userID int) ([]dto.GroupOutput, error) {
	var groups []dto.GroupOutput
	//TODO: pagination !
	result := global.Mdb.Model(&po.Account{}).Where("userID = ? AND type = 1", userID).Find(&groups)
	if result.Error != nil {
		return []dto.GroupOutput{}, result.Error
	}

	return groups, nil
}

// UpdateGroup implements IGroupRepo.
func (gr *GroupRepo) UpdateGroup(userID int, groupInput dto.GroupUpdateInput) error {
	result := global.Mdb.Model(&po.Account{}).Where("userID = ? AND id = ?", userID, groupInput.ID).Select("name", "description").Updates(
		po.Account{
			Name:        groupInput.GroupName,
			Description: groupInput.Description,
		},
	)

	return result.Error
}

func (gr *GroupRepo) CreateGroup(userID int, groupInput dto.GroupCreateInput) error {
	var account = po.Account{
		Type:        1,
		Name:        groupInput.GroupName,
		Description: groupInput.Description,
		UserID:      userID,
	}
	result := global.Mdb.Create(&account)
	if result.Error != nil {
		return result.Error
	}

	percentageRemain := 100
	err := global.Rdb.Set(ctx, fmt.Sprintf("Group-%d", account.ID), percentageRemain, 0).Err()

	return err
}

func NewGroupRepo() IGroupRepo {
	return &GroupRepo{}
}
