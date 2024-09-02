package repo

import (
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/po"
)

type ITransactionRepo interface {
	FindTransaction() ([]po.Transaction, error)
	CreateTransaction(fromAccount dto.Account, toAccount dto.Account, amount int) error
}

type transactionRepo struct{}

// CreateTransaction implements ITransactionRepo.
func (tr *transactionRepo) CreateTransaction(fromAccount dto.Account, toAccount dto.Account, amount int) error {
	panic("unimplemented")
}

// FindTransaction implements ITransactionRepo.
func (tr *transactionRepo) FindTransaction() ([]po.Transaction, error) {
	panic("unimplemented")
}

func NewTransactionRepo() ITransactionRepo {
	return &transactionRepo{}
}
