package dto

import (
	"time"
)

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
	FromAccountID int       `gorm:"column:fromaccountID"`
	ToAccountID   int       `gorm:"column:toaccountID"`
}
