package service

import (
	"errors"

	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/response"
	"gorm.io/gorm"
)

var (
	AccountTypeGroup = 1
)

type ITransactionService interface {
	ListTransaction(userID int) (int, []dto.TransOutput, error)
	MakeTransaction(userID int, transactionInput dto.TransCreateInput) (int, error)
}

type transactionService struct {
	transactionRepo repo.ITransactionRepo
	accountRepo     repo.IAccountRepo
	groupDisRepo    repo.IGroupDisRepo
}

func (ts *transactionService) ListTransaction(userID int) (int, []dto.TransOutput, error) {
	transactions, err := ts.transactionRepo.FindTransaction(userID)
	if err != nil {
		return response.ErrCodeInternal, []dto.TransOutput{}, err
	}

	return response.ErrCodeSuccess, transactions, nil
}

func (ts *transactionService) MakeTransaction(userID int, transactionInput dto.TransCreateInput) (int, error) {
	tx, err := ts.transactionRepo.BeginTransaction()
	if err != nil {
		return response.ErrCodeInternal, err
	}
	defer func() {
		if r := recover(); r != nil {
			ts.transactionRepo.RollbackTransaction(tx)
		}
	}()

	fromAccount, toAccount, errCode, err := ts.validateAccounts(userID, transactionInput)
	if err != nil {
		ts.transactionRepo.RollbackTransaction(tx)
		return errCode, err
	}

	result, err := ts.createTransaction(tx, userID, fromAccount, toAccount, transactionInput)
	if err != nil {
		ts.transactionRepo.RollbackTransaction(tx)
	}

	ts.transactionRepo.CommitTransaction(tx)
	return result, err
}

func (ts *transactionService) validateAccounts(userID int, transactionInput dto.TransCreateInput) (*dto.AccountOutput, *dto.AccountOutput, int, error) {
	var (
		fromAccount, toAccount *dto.AccountOutput
		err                    error
	)

	if transactionInput.FromAccountID == nil && transactionInput.ToAccountID == nil {
		return fromAccount, toAccount, response.ErrCodeInternal, err
	}

	if transactionInput.FromAccountID != nil {
		fromAccount, err = ts.accountRepo.FindAccountByID(userID, *transactionInput.FromAccountID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fromAccount, toAccount, response.ErrCodeAccountNotExist, err
			}
			return fromAccount, toAccount, response.ErrCodeInternal, err
		}
		if fromAccount.Type == AccountTypeGroup {
			return fromAccount, toAccount, response.ErrCodeExternal, nil
		}
	}

	if transactionInput.ToAccountID != nil {
		toAccount, err = ts.accountRepo.FindAccountByID(userID, *transactionInput.ToAccountID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fromAccount, toAccount, response.ErrCodeAccountNotExist, err
			}

			return fromAccount, toAccount, response.ErrCodeInternal, err
		}
	}

	return fromAccount, toAccount, response.ErrCodeSuccess, nil
}

func (ts *transactionService) createTransaction(tx *gorm.DB, userID int, fromAccount, toAccount *dto.AccountOutput, transactionInput dto.TransCreateInput) (int, error) {
	if toAccount != nil {
		err := ts.accountRepo.UpdateAccountBalance(tx, userID, toAccount.ID, toAccount.Balance+transactionInput.Amount)
		if err != nil {
			return response.ErrCodeInternal, err
		}
	}

	if fromAccount != nil {
		if transactionInput.Amount > fromAccount.Balance {
			return response.ErrCodeNotEnoughBalance, errors.New("balance not enough")
		}
		err := ts.accountRepo.UpdateAccountBalance(tx, userID, fromAccount.ID, fromAccount.Balance-transactionInput.Amount)
		if err != nil {
			return response.ErrCodeInternal, err
		}
	}

	err := ts.transactionRepo.CreateTransaction(tx, userID, fromAccount, toAccount, transactionInput)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	if toAccount != nil && toAccount.Type == AccountTypeGroup {
		err := ts.DistributeMoneyToGroup(tx, userID, toAccount.ID, transactionInput)
		if err != nil {
			return response.ErrCodeInternal, err
		}
	}

	return response.ErrCodeSuccess, nil
}

func (ts *transactionService) DistributeMoneyToGroup(tx *gorm.DB, userID, groupID int, transactionInput dto.TransCreateInput) error {
	accounts, err := ts.groupDisRepo.FindAccountByGroupID(userID, groupID)
	if err != nil {
		return err
	}

	presentMoneyInGroup := transactionInput.Amount
	fromAccount, err := ts.accountRepo.FindAccountByIDWithTx(tx, userID, groupID)
	if err != nil {
		return err
	}

	for _, account := range accounts {
		toAccountTemp, err := ts.accountRepo.FindAccountByIDWithTx(tx, userID, account.AccountID)
		if err != nil {
			return err
		}
		toAccount := &dto.AccountOutput{
			ID:          account.Account.ID,
			Type:        account.Account.Type,
			Name:        account.Account.Name,
			Description: account.Account.Description,
			Balance:     toAccountTemp.Balance,
		}

		transaction := dto.TransCreateInput{
			Amount:        presentMoneyInGroup * account.Percentage / 100,
			Description:   transactionInput.Description,
			FromAccountID: &groupID,
			ToAccountID:   &account.AccountID,
		}
		_, err = ts.createTransaction(tx, userID, fromAccount, toAccount, transaction)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewTransactionService(transactionRepo repo.ITransactionRepo, accountRepo repo.IAccountRepo, groupDisRepo repo.IGroupDisRepo) ITransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
		groupDisRepo:    groupDisRepo,
	}
}
