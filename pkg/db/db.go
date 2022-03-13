package db

import (
	"log"
	// "os"

	// "gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	// for postgres
	// dsn := os.Getenv("dsn")

	// db, err := gorm.Open(
	// 	postgres.Open(dsn),
	// 	&gorm.Config{
	// 		Logger: logger.Default.LogMode(logger.Silent),
	// 	},
	// )

	// for sqlite
	db, err := gorm.Open(
		sqlite.Open("godiary.db"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)

	if err != nil {
		log.Fatal("Failed to connect database; \n", err)
		panic(err)
	}

	return db
}
