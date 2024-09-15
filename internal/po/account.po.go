package po

import "gorm.io/gorm"

type Account struct {
	ID          int            `gorm:"primaryKey, column:id; autoIncrement; not null; unique"`
	Type        int            `gorm:"column:type; not null"`
	Name        string         `gorm:"column:name; not null"`
	Description string         `gorm:"column:description"`
	Balance     int            `gorm:"column:balance;default=0"`
	UserID      int            `gorm:"column:userID; not null"`
	User        User           `gorm:"foreignKey:UserID;references:ID"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (t *Account) TableName() string {
	return "go_db_account"
}
