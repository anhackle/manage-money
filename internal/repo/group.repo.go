package repo

import (
	"fmt"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/po"
	"gorm.io/gorm"
)

type IGroupRepo interface {
	CreateGroup(userID int, groupInput dto.GroupCreateInput) error
	FindGroupByID(userID int, groupID int) (dto.GroupOutput, error)
	FindGroupByUserID(userID int, groupInput *dto.GroupListInput) ([]dto.GroupOutput, error)
	UpdateGroup(userID int, groupInput dto.GroupUpdateInput) (*gorm.DB, error)
	DeleteGroup(userID int, groupInput dto.GroupDeleteInput) (*gorm.DB, error)
}

type GroupRepo struct{}

func (gr *GroupRepo) DeleteGroup(userID int, groupInput dto.GroupDeleteInput) (*gorm.DB, error) {
	result := global.Mdb.Model(&po.Account{}).Where("userID = ? AND id = ?", userID, groupInput.ID).Delete(&po.Account{})
	if result.Error != nil {
		return nil, result.Error
	}
	err := global.Rdb.Del(ctx, fmt.Sprintf("Group-%d", groupInput.ID)).Err()
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

func (gr *GroupRepo) FindGroupByID(userID int, groupID int) (dto.GroupOutput, error) {
	var group dto.GroupOutput
	result := global.Mdb.Model(&po.Account{}).Where("userID = ? AND id = ?", userID, groupID).First(&group)
	if result.Error != nil {
		return dto.GroupOutput{}, result.Error
	}

	return group, nil
}

func (gr *GroupRepo) FindGroupByUserID(userID int, groupInput *dto.GroupListInput) ([]dto.GroupOutput, error) {
	var groups []dto.GroupOutput
	result := global.Mdb.Model(&po.Account{}).
		Where("userID = ? AND type = 1", userID).
		Limit(groupInput.PageSize).
		Offset((groupInput.Page - 1) * groupInput.PageSize).
		Find(&groups)
	if result.Error != nil {
		return []dto.GroupOutput{}, result.Error
	}

	return groups, nil
}

func (gr *GroupRepo) UpdateGroup(userID int, groupInput dto.GroupUpdateInput) (*gorm.DB, error) {
	result := global.Mdb.Model(&po.Account{}).Where("userID = ? AND id = ?", userID, groupInput.ID).Select("name", "description").Updates(
		po.Account{
			Name:        groupInput.GroupName,
			Description: groupInput.Description,
		},
	)

	return result, result.Error
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
