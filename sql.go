package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitMySqlByProxy(opts DBProxyOptions) *gorm.DB {

	dbUrl := RegisterDBProxy(opts)

	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		println(err.Error())
	}

	db.Debug()

	sqldb, err := db.DB()
	if err != nil {
		println(err.Error())
	}

	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(10)
	sqldb.SetConnMaxLifetime(time.Hour)

	return db
}
