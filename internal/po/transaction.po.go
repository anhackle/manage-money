package po

import "time"

type Transaction struct {
	ID            int       `gorm:"primaryKey, column:id; autoIncrement; not null; unique"`
	Date          time.Time `gorm:"column:date; not null"`
	Amount        int       `gorm:"column:amount; not null"`
	Description   string    `gorm:"column:desc; not null"`
	FromAccountID *int      `gorm:"column:fromaccountID"`
	ToAccountID   *int      `gorm:"column:toaccountID"`
	UserID        int       `gorm:"column:userID; not null"`
	FromAccount   Account   `gorm:"foreignKey:FromAccountID;references:ID"`
	ToAccount     Account   `gorm:"foreignKey:ToAccountID;references:ID"`
	User          User      `gorm:foreignKey:UserID;references:ID"`
}

func (t *Transaction) TableName() string {
	return "go_db_transaction"
}
