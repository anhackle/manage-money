package repo

import (
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/po"
)

type IAccountRepo interface {
	FindAccount() ([]po.Account, error)
	CreateAccount() error
	UpdateAccount(account dto.Account) error
	DeleteAccount(account dto.Account) error
}

type accountRepo struct{}

// CreateAccount implements IAccountRepo.
func (a *accountRepo) CreateAccount() error {
	panic("unimplemented")
}

// DeleteAccount implements IAccountRepo.
func (a *accountRepo) DeleteAccount(account dto.Account) error {
	panic("unimplemented")
}

// FindAccount implements IAccountRepo.
func (a *accountRepo) FindAccount() ([]po.Account, error) {
	panic("unimplemented")
}

// UpdateAccount implements IAccountRepo.
func (a *accountRepo) UpdateAccount(account dto.Account) error {
	panic("unimplemented")
}

func NewAccountRepo() IAccountRepo {
	return &accountRepo{}
}