package service

import (
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/response"
)

type IAccountService interface {
	ListAccount(userID int) (int, []dto.AccountOutput, error)
	CreateAccount(userID int, account dto.AccountCreateInput) (int, error)
	UpdateAccount(userID int, account dto.AccountUpdateInput) (int, error)
	DeleteAccount(userID int, account dto.AccountDeleteInput) (int, error)
}

type accountService struct {
	accountRepo repo.IAccountRepo
}

// CreateAccount implements IAccountService.
func (as *accountService) CreateAccount(userID int, accountInput dto.AccountCreateInput) (int, error) {
	err := as.accountRepo.CreateAccount(userID, accountInput)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

// ListAccount implements IAccountService.
func (as *accountService) ListAccount(userID int) (int, []dto.AccountOutput, error) {
	accounts, err := as.accountRepo.FindAccount(userID)
	if err != nil {
		return response.ErrCodeInternal, []dto.AccountOutput{}, err
	}

	return response.ErrCodeSuccess, accounts, nil
}

// DeleteAccount implements IAccountService.
func (as *accountService) DeleteAccount(userID int, accountInput dto.AccountDeleteInput) (int, error) {
	err := as.accountRepo.DeleteAccount(userID, accountInput)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

// UpdateAccount implements IAccountService.
func (as *accountService) UpdateAccount(userID int, accountInput dto.AccountUpdateInput) (int, error) {
	err := as.accountRepo.UpdateAccount(userID, accountInput)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

func NewAccountService(accountRepo repo.IAccountRepo) IAccountService {
	return &accountService{
		accountRepo: accountRepo,
	}
}
