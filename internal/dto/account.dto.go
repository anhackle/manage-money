package dto

type AccountOutput struct {
	ID          int    `json:"accountID" gorm:"column:id"`
	Type        int    `json:"type" gorm:"column:type"`
	Name        string `json:"Name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	Balance     int    `json:"balance" gorm:"column:balance"`
}

type AccountCreateInput struct {
	AccountName string `json:"accountName" binding:"required,ascii,max=50"`
	Description string `json:"description" binding:"required,ascii,max=255"`
}

type AccountUpdateInput struct {
	ID          int    `json:"accountID" binding:"required,number,min=0"`
	AccountName string `json:"accountName" binding:"required,ascii,max=50"`
	Description string `json:"description" binding:"required,ascii,max=255"`
}

type AccountDeleteInput struct {
	ID int `json:"accountID" binding:"required,number"`
}
