package dto

type GroupDisListInput struct {
	GroupID int `json:"groupID" binding:"required,number,min=0"`
}

type GroupDisCreateInput struct {
	GroupID    int `json:"groupID" binding:"required,number,min=0"`
	AccountID  int `json:"accountID" binding:"required,number,min=0"`
	Percentage int `json:"percentage" binding:"required,number,min=0"`
}

type GroupDisDeleteInput struct {
	GroupID   int `json:"groupID" binding:"required,number,min=0"`
	AccountID int `json:"accountID" binding:"required,number,min=0"`
}

type GroupDisOutput struct {
	ID         int `gorm:"column:id"`
	GroupID    int `gorm:"column:groupID"`
	AccountID  int `gorm:"column:accountID"`
	Percentage int `gorm:"column:percentage"`
}
