package po

import "gorm.io/gorm"

type GroupDistributed struct {
	ID         int            `gorm:"primaryKey, column:id; not null; autoIncrement; unique"`
	UserID     int            `gorm:"column:userID; not null"`
	GroupID    int            `gorm:"column:groupID; not null"`
	AccountID  int            `gorm:"column:accountID; not null"`
	Percentage int            `gorm:"column:percentage; not null"`
	Group      Account        `gorm:"foreignKey:GroupID; references:ID"`
	Account    Account        `gorm:"foreignKey:AccountID; references:ID"`
	User       User           `grom:"foreignkey:UserID; references:ID"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (t *GroupDistributed) TableName() string {
	return "go_db_groupdistributed"
}
