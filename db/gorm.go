package db

import (
	"fmt"
	"log"
	"os"
	"time"
	"todo-list-api/config"
	"todo-list-api/exception"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormConnection(conf config.Config) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // Slow SQL threshold
			LogLevel:      logger.Error, // Log level    // Ignore ErrRecordNotFound error for logger
			Colorful:      false,        // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Get("MYSQL_USER"), conf.Get("MYSQL_PASSWORD"), conf.Get("MYSQL_HOST"), conf.Get("MYSQL_DBNAME"))), &gorm.Config{
		Logger: newLogger,
	})
	exception.PanicWhenError(err)

	return db
}
