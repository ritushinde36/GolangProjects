package pkg

import (
	"jwt-auth-sql/models"
)

func CreateTable() {
	DB.AutoMigrate(&models.User{})
}
