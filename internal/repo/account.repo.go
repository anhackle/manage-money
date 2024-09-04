package repo

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/po"
)

type IAccountRepo interface {
	UpdateAccountBalance(userID, accountID, balance int) error
	FindAccountByUserID(userID int) ([]dto.AccountOutput, error)
	FindAccountByID(userID int, accountID int) (dto.AccountOutput, error)
	CreateAccount(userID int, accountInput dto.AccountCreateInput) error
	UpdateAccount(userID int, accountInput dto.AccountUpdateInput) error
	DeleteAccount(userID int, accountInput dto.AccountDeleteInput) error
}

type accountRepo struct{}

func (ar *accountRepo) UpdateAccountBalance(userID, accountID, balance int) error {
	result := global.Mdb.Table("go_db_account").Where("userID = ? AND id = ?", userID, accountID).Updates(
		po.Account{
			Balance: balance,
		},
	)

	return result.Error
}

func (ar *accountRepo) FindAccountByID(userID int, accountID int) (dto.AccountOutput, error) {
	var account dto.AccountOutput
	result := global.Mdb.Table("go_db_account").Where("userID = ? AND id = ?", userID, accountID).First(&account)
	if result.Error != nil {
		return dto.AccountOutput{}, result.Error
	}

	return account, nil
}

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

func (ar *accountRepo) FindAccountByUserID(userID int) ([]dto.AccountOutput, error) {
	var accounts []dto.AccountOutput
	//TODO: pagination !
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
