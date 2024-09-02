package repo

import (
	"github.com/anle/codebase/internal/po"
)

type ITransactionRepo interface {
	FindTransaction() ([]po.Transaction, error)
	CreateTransaction() error
}

type transactionRepo struct{}

// CreateTransaction implements ITransactionRepo.
func (tr *transactionRepo) CreateTransaction() error {
	panic("unimplemented")
}

// FindTransaction implements ITransactionRepo.
func (tr *transactionRepo) FindTransaction() ([]po.Transaction, error) {
	panic("unimplemented")
}

func NewTransactionRepo() ITransactionRepo {
	return &transactionRepo{}
}
