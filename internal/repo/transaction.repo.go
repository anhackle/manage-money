package repo

import (
	"time"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/dto"
	"github.com/anle/codebase/internal/po"
	"gorm.io/gorm"
)

type ITransactionRepo interface {
	BeginTransaction() (*gorm.DB, error)
	CommitTransaction(tx *gorm.DB) error
	RollbackTransaction(tx *gorm.DB) error
	FindTransaction(userID int, transactionInput *dto.TransListInput) ([]dto.TransOutput, error)
	FindTransactionByUserIDAndCurrency(tx *gorm.DB, userID, currencyID, fromAccountID int) ([]dto.TransOutput, error)
	CreateTransaction(tx *gorm.DB, userID int, fromAccount, toAccount *dto.AccountOutput, transactionInput dto.TransCreateInput) error
}

type transactionRepo struct {
}

func (tr *transactionRepo) BeginTransaction() (*gorm.DB, error) {
	tx := global.Mdb.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (tr *transactionRepo) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (tr *transactionRepo) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (tr *transactionRepo) CreateTransaction(tx *gorm.DB, userID int, fromAccount, toAccount *dto.AccountOutput, transactionInput dto.TransCreateInput) error {
	var transaction = po.Transaction{
		Date:          time.Now(),
		Amount:        transactionInput.Amount,
		CurrencyID:    transactionInput.CurrencyID,
		Description:   transactionInput.Description,
		FromAccountID: transactionInput.FromAccountID,
		ToAccountID:   transactionInput.ToAccountID,
		UserID:        userID,
	}
	result := tx.Create(&transaction)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ts *transactionRepo) FindTransaction(userID int, transactionInput *dto.TransListInput) ([]dto.TransOutput, error) {
	var transactions []dto.TransOutput
	result := global.Mdb.
		Model(&po.Transaction{}).
		Where("userID = ?", userID).
		Limit(transactionInput.PageSize).
		Offset((transactionInput.Page - 1) * transactionInput.PageSize).
		Find(&transactions)
	if result.Error != nil {
		return []dto.TransOutput{}, result.Error
	}

	return transactions, nil
}

func (ts *transactionRepo) FindTransactionByUserIDAndCurrency(tx *gorm.DB, userID, currencyID, fromAccountID int) ([]dto.TransOutput, error) {
	var transactions []dto.TransOutput
	result := tx.
		Model(&po.Transaction{}).
		Where("userID = ? AND currencyID = ? AND (fromAccountID = ? OR toAccountID = ?)", userID, currencyID, fromAccountID, fromAccountID).
		Find(&transactions)
	if result.Error != nil {
		return []dto.TransOutput{}, result.Error
	}

	return transactions, nil
}

func NewTransactionRepo() ITransactionRepo {
	return &transactionRepo{}
}
