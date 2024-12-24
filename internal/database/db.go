package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

func Init(dsn string) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)                 // Maximum number of open connections to the database
	sqlDB.SetMaxIdleConns(10)                  // Maximum number of idle connections in the pool
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Maximum amount of time a connection may be reused
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // Maximum amount of time an idle connection is kept in the pool
}
