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
	FindTransaction(userID int) ([]dto.TransOutput, error)
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

func (ts *transactionRepo) FindTransaction(userID int) ([]dto.TransOutput, error) {
	var transactions []dto.TransOutput
	//TODO: pagination !
	result := global.Mdb.Model(&po.Transaction{}).Where("userID = ?", userID).Find(&transactions)
	if result.Error != nil {
		return []dto.TransOutput{}, result.Error
	}

	return transactions, nil
}

func NewTransactionRepo() ITransactionRepo {
	return &transactionRepo{}
}
