package dto

import (
	"time"
)

type TransListInput struct {
	Page     int `json:"page" binding:"required,number,gt=0"`
	PageSize int `json:"pagesize" binding:"required,number,gt=0,lt=50"`
}

type TransCreateInput struct {
	Amount        int    `json:"amount" binding:"required,number,gt=0"`
	Description   string `json:"description"`
	FromAccountID *int   `json:"fromAccountID"`
	ToAccountID   *int   `json:"toAccountID"`
}

type TransOutput struct {
	Date          time.Time `gorm:"column:date; not null"`
	Amount        int       `gorm:"column:amount; not null"`
	Description   string    `gorm:"column:desc; not null"`
	FromAccountID *int      `gorm:"column:fromaccountID"`
	ToAccountID   int       `gorm:"column:toaccountID"`
}
