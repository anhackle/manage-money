package service

import (
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/po"
	"github.com/anle/codebase/internal/repo"
)

type ITransactionService interface {
	ListTransaction() ([]po.Token, error)
	CreateTransaction(fromAccount dto.Account, toAccount dto.Account, amount int) error
}

type transactionService struct {
	transactionRepo repo.ITransactionRepo
}

// ListTransaction implements ITransactionService.
func (ts *transactionService) ListTransaction() ([]po.Token, error) {
	panic("unimplemented")
}

// MakeTransaction implements ITransactionService.
func (ts *transactionService) CreateTransaction(fromAccount dto.Account, toAccount dto.Account, amount int) error {
	panic("unimplemented")
}

func NewTransactionService(transactionRepo repo.ITransactionRepo) ITransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
	}
}
