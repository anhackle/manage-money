package dto

type GroupOutput struct {
	ID          int    `json:"groupID" gorm:"column:id"`
	GroupName   string `json:"groupName" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

type GroupListInput struct {
	Page     int `json:"page" binding:"required,number,gt=0"`
	PageSize int `json:"pagesize" binding:"required,number,gt=0,lt=50"`
}

type GroupCreateInput struct {
	GroupName   string `json:"groupName" binding:"required,ascii,max=50"`
	Description string `json:"description" binding:"required,ascii,max=255"`
}

type GroupUpdateInput struct {
	ID          int    `json:"groupID" binding:"required,number,min=0"`
	GroupName   string `json:"groupName" binding:"required,ascii,max=50"`
	Description string `json:"description" binding:"required,ascii,max=255"`
}

type GroupDeleteInput struct {
	ID int `json:"groupID" binding:"required,number"`
}
