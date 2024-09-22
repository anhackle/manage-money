package repo

import (
	"fmt"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/po"
)

type IGroupDisRepo interface {
	Create(userID int, percentageRemain int, groupDisInput dto.GroupDisCreateInput) error
	Delete(userID int, percentageRemain int, groupDisInput dto.GroupDisDeleteInput) error
	FindAccountByGroupID(userID, groupID int) ([]po.GroupDistributed, error)
	FindGroupDisByGroupAccount(userID int, groupDisInput dto.GroupDisDeleteInput) (dto.GroupDisOutput, error)
	FindGroupDisByGroupID(userID int, groupDisInput dto.GroupDisListInput) ([]dto.GroupDisOutput, error)
	FindPercentageRemain(userID, groupID int) (int, error)
}

type groupDisRepo struct{}

func (gdr *groupDisRepo) FindPercentageRemain(userID, groupID int) (int, error) {
	var percentageRemain int
	percentageRemain, err := global.Rdb.Get(ctx, fmt.Sprintf("Group-%d", groupID)).Int()
	if err != nil {
		err = global.Mdb.Model(&po.GroupDistributed{}).
			Select("100 - IFNULL(SUM(percentage), 0)").
			Where("userID = ? AND groupID = ?", userID, groupID).
			Scan(&percentageRemain).Error
		if err != nil {
			return -1, err
		}

		global.Rdb.Set(ctx, fmt.Sprintf("Group-%d", groupID), percentageRemain, 0).Err()

		return percentageRemain, nil
	}

	return percentageRemain, nil
}

func (gdr *groupDisRepo) FindAccountByGroupID(userID, groupID int) ([]po.GroupDistributed, error) {
	var accounts []po.GroupDistributed
	result := global.Mdb.Model(&po.GroupDistributed{}).
		Preload("Account").
		Where("userID = ? AND groupID = ?", userID, groupID).
		Find(&accounts)
	if result.Error != nil {
		return []po.GroupDistributed{}, result.Error
	}

	return accounts, nil
}

func (gdr *groupDisRepo) FindGroupDisByGroupAccount(userID int, groupDisInput dto.GroupDisDeleteInput) (dto.GroupDisOutput, error) {
	var groupDis dto.GroupDisOutput
	result := global.Mdb.Model(&po.GroupDistributed{}).Where("userID = ? AND groupID = ? AND accountID = ?", userID, groupDisInput.GroupID, groupDisInput.AccountID).First(&groupDis)
	if result.Error != nil {
		return dto.GroupDisOutput{}, result.Error
	}

	return groupDis, nil
}

func (gdr *groupDisRepo) FindGroupDisByGroupID(userID int, groupDisInput dto.GroupDisListInput) ([]dto.GroupDisOutput, error) {
	var groupDis []dto.GroupDisOutput
	result := global.Mdb.Model(&po.GroupDistributed{}).Where("userID = ? AND groupID = ?", userID, groupDisInput.GroupID).Find(&groupDis)
	if result.Error != nil {
		return []dto.GroupDisOutput{}, result.Error
	}

	return groupDis, nil
}

func (gdr *groupDisRepo) Delete(userID int, percentageRemain int, groupDisInput dto.GroupDisDeleteInput) error {
	groupDis, err := gdr.FindGroupDisByGroupAccount(userID, groupDisInput)
	if err != nil {
		return err
	}

	result := global.Mdb.Model(&po.GroupDistributed{}).Where(
		"userID = ? AND groupID = ? AND accountID = ?",
		userID,
		groupDisInput.GroupID,
		groupDisInput.AccountID).Delete(
		&po.GroupDistributed{},
	)
	if result.Error != nil {
		return result.Error
	}

	err = global.Rdb.Set(ctx, fmt.Sprintf("Group-%d", groupDisInput.GroupID), percentageRemain+groupDis.Percentage, 0).Err()

	return err
}

func (gdr *groupDisRepo) Create(userID int, percentageRemain int, groupDisInput dto.GroupDisCreateInput) error {
	var groupDis = po.GroupDistributed{
		UserID:     userID,
		GroupID:    groupDisInput.GroupID,
		AccountID:  groupDisInput.AccountID,
		Percentage: groupDisInput.Percentage,
	}
	result := global.Mdb.Create(&groupDis)
	if result.Error != nil {
		return result.Error
	}

	err := global.Rdb.Set(ctx, fmt.Sprintf("Group-%d", groupDisInput.GroupID), percentageRemain-groupDis.Percentage, 0).Err()

	return err
}

func NewGroupDisRepo() IGroupDisRepo {
	return &groupDisRepo{}
}
