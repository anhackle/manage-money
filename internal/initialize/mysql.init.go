package initialize

import (
	"fmt"
	"time"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DbName)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	checkErrorPanic(err, "InitMysql initialization error")

	global.Logger.Info("Initializing MySQL Successfully")

	global.Mdb = db
	SetPool()
	MigrateTable()
}

func SetPool() {
	m := global.Config.Mysql

	sqlDB, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("MySQL error: %s::", err)
	}

	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))

}

func MigrateTable() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Token{},
	)
	if err != nil {
		fmt.Println("Migrating tables error", err)
	}
}
