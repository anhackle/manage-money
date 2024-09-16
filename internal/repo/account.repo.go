package repo

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/po"
	"gorm.io/gorm"
)

type IAccountRepo interface {
	UpdateAccountBalance(tx *gorm.DB, userID, accountID, balance int) error
	FindAccountByUserID(userID int, accountInput *dto.AccountListInput) ([]dto.AccountOutput, error)
	FindAccountByID(userID int, accountID int) (*dto.AccountOutput, error)
	FindAccountByIDWithTx(tx *gorm.DB, userID int, accountID int) (*dto.AccountOutput, error)
	CreateAccount(userID int, accountInput dto.AccountCreateInput) error
	UpdateAccount(userID int, accountInput dto.AccountUpdateInput) (*gorm.DB, error)
	DeleteAccount(userID int, accountInput dto.AccountDeleteInput) (*gorm.DB, error)
}

type accountRepo struct{}

func (ar *accountRepo) UpdateAccountBalance(tx *gorm.DB, userID, accountID, balance int) error {
	result := tx.Model(&po.Account{}).Where("userID = ? AND id = ?", userID, accountID).Select("balance").Updates(
		po.Account{
			Balance: balance,
		},
	)

	return result.Error
}

func (ar *accountRepo) FindAccountByID(userID int, accountID int) (*dto.AccountOutput, error) {
	var account *dto.AccountOutput
	result := global.Mdb.Model(&po.Account{}).Where("userID = ? AND id = ?", userID, accountID).First(&account)
	if result.Error != nil {
		return &dto.AccountOutput{}, result.Error
	}

	return account, nil
}

func (ar *accountRepo) FindAccountByIDWithTx(tx *gorm.DB, userID int, accountID int) (*dto.AccountOutput, error) {
	var account *dto.AccountOutput
	result := tx.Model(&po.Account{}).Where("userID = ? AND id = ?", userID, accountID).First(&account)
	if result.Error != nil {
		return &dto.AccountOutput{}, result.Error
	}

	return account, nil
}

func (ar *accountRepo) CreateAccount(userID int, accountInput dto.AccountCreateInput) error {
	var account = po.Account{
		Type:        0,
		Name:        accountInput.AccountName,
		Description: accountInput.Description,
		UserID:      userID,
	}
	result := global.Mdb.Create(&account)

	return result.Error
}

func (ar *accountRepo) FindAccountByUserID(userID int, accountInput *dto.AccountListInput) ([]dto.AccountOutput, error) {
	var accounts []dto.AccountOutput
	result := global.Mdb.Model(&po.Account{}).
		Where("userID = ? AND type = 0", userID).
		Limit(accountInput.PageSize).
		Offset((accountInput.Page - 1) * accountInput.PageSize).
		Find(&accounts)
	if result.Error != nil {
		return []dto.AccountOutput{}, result.Error
	}

	return accounts, nil
}

func (ar *accountRepo) DeleteAccount(userID int, accountInput dto.AccountDeleteInput) (*gorm.DB, error) {
	result := global.Mdb.Model(&po.Account{}).Where("userID = ? AND id = ?", userID, accountInput.ID).Delete(&po.Account{})
	return result, result.Error
}

func (ar *accountRepo) UpdateAccount(userID int, accountInput dto.AccountUpdateInput) (*gorm.DB, error) {
	result := global.Mdb.Model(&po.Account{}).Where("userID = ? AND id = ?", userID, accountInput.ID).Select("name", "description").Updates(
		po.Account{
			Name:        accountInput.AccountName,
			Description: accountInput.Description,
		},
	)

	return result, result.Error
}

func NewAccountRepo() IAccountRepo {
	return &accountRepo{}
}
