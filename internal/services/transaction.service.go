package service

import (
	"errors"

	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/response"
	"gorm.io/gorm"
)

type ITransactionService interface {
	ListTransaction(userID int) (int, []dto.TransOutput, error)
	CreateTransaction(userID int, transactionInput dto.TransCreateInput) (int, error)
}

type transactionService struct {
	transactionRepo repo.ITransactionRepo
	accountRepo     repo.IAccountRepo
}

func (ts *transactionService) ListTransaction(userID int) (int, []dto.TransOutput, error) {
	transactions, err := ts.transactionRepo.FindTransaction(userID)
	if err != nil {
		return response.ErrCodeInternal, []dto.TransOutput{}, err
	}

	return response.ErrCodeSuccess, transactions, nil
}

func (ts *transactionService) CreateTransaction(userID int, transactionInput dto.TransCreateInput) (int, error) {
	if transactionInput.FromAccountID == nil && transactionInput.ToAccountID == nil {
		return response.ErrCodeExternal, nil
	}

	var (
		fromAccount, toAccount dto.AccountOutput
		err                    error
	)

	if transactionInput.FromAccountID != nil {
		fromAccount, err = ts.accountRepo.FindAccountByID(userID, *transactionInput.FromAccountID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return response.ErrCodeAccountNotExist, nil
			}

			return response.ErrCodeInternal, nil
		}
	}

	if transactionInput.ToAccountID != nil {
		toAccount, err = ts.accountRepo.FindAccountByID(userID, *transactionInput.ToAccountID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return response.ErrCodeAccountNotExist, nil
			}

			return response.ErrCodeInternal, nil
		}
	}

	if transactionInput.ToAccountID == nil {
		err = ts.transactionRepo.CreateTransactionNoToAccount(userID, fromAccount, transactionInput)
		if err != nil {
			return response.ErrCodeInternal, err
		}

		return response.ErrCodeSuccess, nil
	}

	if transactionInput.FromAccountID == nil {
		err = ts.transactionRepo.CreateTransactionNoFromAccount(userID, toAccount, transactionInput)
		if err != nil {
			return response.ErrCodeInternal, err
		}

		return response.ErrCodeSuccess, nil
	}

	if transactionInput.Amount > fromAccount.Balance {
		return response.ErrCodeNotEnoughBalance, nil
	}

	err = ts.transactionRepo.CreateTransaction(userID, fromAccount, toAccount, transactionInput)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

func NewTransactionService(transactionRepo repo.ITransactionRepo, accountRepo repo.IAccountRepo) ITransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
	}
}
