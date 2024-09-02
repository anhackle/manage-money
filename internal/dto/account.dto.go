package dto

type Account struct {
	ID          int    `json:"accountID" binding:"required,number"`
	AccountName string `json:"accountName" binding:"ascii,max=50"`
	Description string `json:"description" binding:"ascii,max=255"`
}
