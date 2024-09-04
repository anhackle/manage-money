package repo

import (
	"time"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/po"
)

type ITransactionRepo interface {
	FindTransaction(userID int) ([]dto.TransOutput, error)
	CreateTransaction(userID int, fromAccount, toAccount dto.AccountOutput, transactionInput dto.TransCreateInput) error
	CreateTransactionNoFromAccount(userID int, toAccount dto.AccountOutput, transactionInput dto.TransCreateInput) error
	CreateTransactionNoToAccount(userID int, fromAccount dto.AccountOutput, transactionInput dto.TransCreateInput) error
}

type transactionRepo struct {
	accountRepo IAccountRepo
}

func (tr *transactionRepo) CreateTransactionNoFromAccount(userID int, toAccount dto.AccountOutput, transactionInput dto.TransCreateInput) error {
	tx := global.Mdb.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Add toAccount
	err := tr.accountRepo.UpdateAccountBalance(userID, *transactionInput.ToAccountID, toAccount.Balance+transactionInput.Amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	//Create transaction
	var transaction = po.Transaction{
		Date:          time.Now(),
		Amount:        transactionInput.Amount,
		Description:   transactionInput.Description,
		FromAccountID: nil,
		ToAccountID:   transactionInput.ToAccountID,
		UserID:        userID,
	}
	result := global.Mdb.Create(&transaction)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()
	return nil
}

func (tr *transactionRepo) CreateTransactionNoToAccount(userID int, fromAccount dto.AccountOutput, transactionInput dto.TransCreateInput) error {
	tx := global.Mdb.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Deduct fromAccount
	err := tr.accountRepo.UpdateAccountBalance(userID, *transactionInput.FromAccountID, fromAccount.Balance-transactionInput.Amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	//Create transaction
	var transaction = po.Transaction{
		Date:          time.Now(),
		Amount:        transactionInput.Amount,
		Description:   transactionInput.Description,
		FromAccountID: transactionInput.FromAccountID,
		ToAccountID:   nil,
		UserID:        userID,
	}
	result := global.Mdb.Create(&transaction)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()
	return nil
}

func (tr *transactionRepo) CreateTransaction(userID int, fromAccount, toAccount dto.AccountOutput, transactionInput dto.TransCreateInput) error {
	tx := global.Mdb.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//Deduct fromAccount
	err := tr.accountRepo.UpdateAccountBalance(userID, *transactionInput.FromAccountID, fromAccount.Balance-transactionInput.Amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	//Add toAccount
	err = tr.accountRepo.UpdateAccountBalance(userID, *transactionInput.ToAccountID, toAccount.Balance+transactionInput.Amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	//Create transaction
	var transaction = po.Transaction{
		Date:          time.Now(),
		Amount:        transactionInput.Amount,
		Description:   transactionInput.Description,
		FromAccountID: transactionInput.FromAccountID,
		ToAccountID:   transactionInput.ToAccountID,
		UserID:        userID,
	}
	result := global.Mdb.Create(&transaction)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()
	return nil
}

func (ts *transactionRepo) FindTransaction(userID int) ([]dto.TransOutput, error) {
	var transactions []dto.TransOutput
	//TODO: pagination !
	result := global.Mdb.Table("go_db_transaction").Where("userID = ?", userID).Find(&transactions)
	if result.Error != nil {
		return []dto.TransOutput{}, result.Error
	}

	return transactions, nil
}

func NewTransactionRepo(accountRepo IAccountRepo) ITransactionRepo {
	return &transactionRepo{
		accountRepo: accountRepo,
	}
}
