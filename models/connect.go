package models

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Tabler interface {
	TableName() string
}

var DB *gorm.DB

func ConnectDb() {
	logLevel := logger.Info
	if gin.Mode() == gin.ReleaseMode {
		logLevel = logger.Silent
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)
	db, err := gorm.Open(sqlite.Open("segments.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	if 1 == 0 {
		db.AutoMigrate(
			&ColorType{},
			&Color{},
			&Company{},
			&Section{},
			&Rack{},
			&OrderNumber{},
			&Segment{},
			&User{},
		)
	}
	DB = db
}
