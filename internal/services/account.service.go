package service

import (
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/po"
	"github.com/anle/codebase/internal/repo"
)

type IAccountService interface {
	ListAccount() (po.Account, error)
	CreateAccount() error
	UpdateAccount(dto.Account) error
	DeleteAccount(dto.Account) error
}

type accountService struct {
	accountRepo repo.IAccountRepo
}

// CreateAccount implements IAccountService.
func (a *accountService) CreateAccount() error {
	panic("unimplemented")
}

// DeleteAccount implements IAccountService.
func (a *accountService) DeleteAccount(dto.Account) error {
	panic("unimplemented")
}

// ListAccount implements IAccountService.
func (a *accountService) ListAccount() (po.Account, error) {
	panic("unimplemented")
}

// UpdateAccount implements IAccountService.
func (a *accountService) UpdateAccount(dto.Account) error {
	panic("unimplemented")
}

func NewAccountService(accountRepo repo.IAccountRepo) IAccountService {
	return &accountService{
		accountRepo: accountRepo,
	}
}
