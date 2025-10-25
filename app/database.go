package app

import (
	"log"

	"github.com/sorfian/go-todo-list/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	// Load config
	config := LoadConfig()

	// Set log level based on environment
	logLevel := logger.Info
	if config.AppEnv == "production" {
		logLevel = logger.Error
	}

	// Open a database connection
	dialect := mysql.Open(config.GetDSN())
	gormDB, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	helper.PanicIfError(err)

	// Configure a connection pool
	db, err := gormDB.DB()
	helper.PanicIfError(err)
	db.SetMaxIdleConns(config.Database.MaxIdleConns)
	db.SetMaxOpenConns(config.Database.MaxOpenConns)
	db.SetConnMaxLifetime(config.Database.ConnMaxLifetime)
	db.SetConnMaxIdleTime(config.Database.ConnMaxIdleTime)

	log.Printf("Database connected successfully to %s:%s/%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	return gormDB
}
