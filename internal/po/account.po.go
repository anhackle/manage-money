package po

type Account struct {
	ID          int    `gorm:"primaryKey, column:id; autoIncrement; not null; unique"`
	AccountName string `gorm:"column:accountName; not null"`
	Description string `gorm:"column:description"`
	Balance     int    `gorm:"column:balance;default=0"`
	UserID      int    `gorm:"column:userID; not null"`
	User        User   `gorm:"foreignKey:UserID"`
}

func (t *Account) TableName() string {
	return "go_db_account"
}
