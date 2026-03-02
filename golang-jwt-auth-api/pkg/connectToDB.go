package pkg

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB = ConnectToDB()

func ConnectToDB() *gorm.DB {
	LoadEnvVariables()

	dsn_name := os.Getenv("DSN")

	db, err := gorm.Open(mysql.Open(dsn_name), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db

}
