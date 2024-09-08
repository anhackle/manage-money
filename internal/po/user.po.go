package po

type User struct {
	ID       int    `gorm:"primaryKey, column:id; autoIncrement; not null; unique;"`
	Email    string `gorm:"column:email; size:50; not null; unique" json:"email" binding:"required,email,lowercase,max=50"`
	Password string `gorm:"column:password; not null" json:"password" binding:"required,password,min=8,max=20"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
