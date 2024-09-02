package repo

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/po"
)

type IAccountRepo interface {
	FindAccount(userID int) ([]dto.AccountOutput, error)
	CreateAccount(userID int, accountInput dto.AccountCreateInput) error
	// UpdateAccount(accountInput dto.Account) error
	// DeleteAccount(accountInput dto.Account) error
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

// FindAccount implements IAccountRepo.
func (ar *accountRepo) FindAccount(userID int) ([]dto.AccountOutput, error) {
	var accounts []dto.AccountOutput
	result := global.Mdb.Table("go_db_account").Where("userID = ?", userID).Find(&accounts)
	if result.Error != nil {
		return []dto.AccountOutput{}, result.Error
	}

	return accounts, nil
}

// func (ar *accountRepo) DeleteAccount(account dto.Account) error {
// 	panic("unimplemented")
// }

// func (ar *accountRepo) UpdateAccount(account dto.Account) error {
// 	panic("unimplemented")
// }

func NewAccountRepo() IAccountRepo {
	return &accountRepo{}
}
