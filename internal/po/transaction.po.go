package po

import "time"

type Transaction struct {
	ID            int       `gorm:"primaryKey, column:id; autoIncrement; not null; unique"`
	Date          time.Time `gorm:"column:date; not null"`
	Description   string    `gorm:"column:desc; not null"`
	FromAccountID int       `gorm:"fromaccountID; not null"`
	ToAccountID   int       `gorm:"toaccountID; not null"`
	FromAccount   Account   `gorm:"foreignKey:FromAccountID"`
	ToAccount     Account   `gorm:"foreignKey:ToAccountID"`
}

func (t *Transaction) TableName() string {
	return "go_db_transaction"
}
