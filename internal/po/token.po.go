package po

type Token struct {
	ID     int    `gorm:"primaryKey; column:id; autoIncrement; not null; unique;"`
	Token  string `gorm:"column:token; size:60; not null"`
	UserID int    `gorm:"column:userid; not null"`
	User   User   `gorm:"foreignKey:UserID;references:ID"`
}

func (t *Token) TableName() string {
	return "go_db_token"
}
