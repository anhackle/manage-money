package repo

import "github.com/anle/codebase/internal/po"

type ITransactionRepo interface {
	FindTransaction() ([]po.Transaction, error)
}

type transactionRepo struct{}

// FindTransaction implements ITransactionRepo.
func (t *transactionRepo) FindTransaction() ([]po.Transaction, error) {
	panic("unimplemented")
}

func NewTransactionRepo() ITransactionRepo {
	return &transactionRepo{}
}
