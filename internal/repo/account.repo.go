package repo

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/po"
)

type IAccountRepo interface {
	FindAccount(userID int) ([]dto.AccountOutput, error)
	CreateAccount(userID int, accountInput dto.AccountCreateInput) error
	UpdateAccount(userID int, accountInput dto.AccountUpdateInput) error
	DeleteAccount(userID int, accountInput dto.AccountDeleteInput) error
}

type accountRepo struct{}

// CreateAccount implements IAccountRepo.
func (ar *accountRepo) CreateAccount(userID int, accountInput dto.AccountCreateInput) error {
	var account = po.Account{
		AccountName: accountInput.AccountName,
		Description: accountInput.Description,
		UserID:      userID,
	}
	result := global.Mdb.Create(&account)

	return result.Error
}

func (ar *accountRepo) FindAccount(userID int) ([]dto.AccountOutput, error) {
	var accounts []dto.AccountOutput
	result := global.Mdb.Table("go_db_account").Where("userID = ?", userID).Find(&accounts)
	if result.Error != nil {
		return []dto.AccountOutput{}, result.Error
	}

	return accounts, nil
}

func (ar *accountRepo) DeleteAccount(userID int, accountInput dto.AccountDeleteInput) error {
	result := global.Mdb.Table("go_db_account").Where("userID = ? AND id = ?", userID, accountInput.ID).Delete(&po.Account{})
	return result.Error
}

func (ar *accountRepo) UpdateAccount(userID int, accountInput dto.AccountUpdateInput) error {
	result := global.Mdb.Table("go_db_account").Where("userID = ? AND id = ?", userID, accountInput.ID).Updates(
		po.Account{
			AccountName: accountInput.AccountName,
			Description: accountInput.Description,
		},
	)

	return result.Error
}

func NewAccountRepo() IAccountRepo {
	return &accountRepo{}
}
