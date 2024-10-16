package initialization

import (
	"backend-go/global"
	"backend-go/internal/po"
	"fmt"

	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CheckPanicError(err error, errorMsg string) {
	if err != nil {
		global.Logger.Error(errorMsg, zap.Error(err)) // write error to log
		panic(err)
	}
}

func InitMysql() {
	mysqlConfig := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(dsn, mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.
	Port, mysqlConfig.DbName)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	CheckPanicError(err, "Initialize mysql failed")
	global.Logger.Info("Initialize mysql successfully")
	global.MysqlDB = db

	// set pool
	SetPool()
	migrateTables()
}

func SetPool() {
	sqlDb, err := global.MysqlDB.DB()

	if err != nil {
		fmt.Printf("mysql error: %s::", err)
	}

	sqlDb.SetConnMaxIdleTime(time.Duration(global.Config.Mysql.MaxIdleConns))
	sqlDb.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(global.Config.Mysql.ConnMaxLifetime))
}

func migrateTables() {
	err := global.MysqlDB.AutoMigrate(&po.User{}, &po.Role{})
	if err != nil {
		fmt.Println("Migrate tables error: ", err)
	}
}
