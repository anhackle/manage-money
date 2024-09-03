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
	Date          time.Time     `json:"date"`
	Amount        int           `json:"amount"`
	Description   string        `json:"description"`
	FromAccountID AccountOutput `json:"fromAccount"`
	ToAccountID   AccountOutput `json:"toACcount"`
}
