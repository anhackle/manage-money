package po

type User struct {
	ID       int    `gorm:"column:id; autoIncrement; not null; unique;"`
	Email    string `gorm:"column:email; size:50; not null" json:"email" binding:"required"`
	Password string `gorm:"column:password; not null" json:"password" binding:"required"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
