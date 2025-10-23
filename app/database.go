package app

import (
	"time"

	"github.com/sorfian/go-todo-list/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	dialect := mysql.Open("root@tcp(localhost:3306)/go_todo_list?charset=utf8mb4&parseTime=True&loc=Local")
	gormDB, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	helper.PanicIfError(err)

	db, err := gormDB.DB()
	helper.PanicIfError(err)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return gormDB
}
