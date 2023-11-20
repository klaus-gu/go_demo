package operate

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var MyDB = Connect("root", "1qaz2wsx3edc..", "127.0.0.1:3306", "article")

/**
 * 拉取依赖：go get -u gorm.io/driver/mysql ，go get -u gorm.io/gorm
 */
func Connect(username, password, addr, dbName string) *gorm.DB {
	dsn := username + ":" + password + "@tcp(" + addr + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库建立连接异常：", err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("数据库连接异常：", err.Error())
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
