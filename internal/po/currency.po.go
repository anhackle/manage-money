package po

type Currency struct {
	ID           int    `gorm:"primaryKey, column:id; autoIncrement; not null; unique"`
	Name         string `gorm:"column:name; not null"`
	ExchangeRate int    `gorm:"column:exchange_rate; not null"`
}

func (c *Currency) TableName() string {
	return "go_db_currency"
}
