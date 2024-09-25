package models

import (
	"log"

	"github.com/braam76/auth-backend/api/v1/database/mysql"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Username string
	Password string
}

func init() {
	err := mysql.DB.AutoMigrate(User{})
	if err != nil {
		log.Fatalf("Error while migrating User table: %s", err)
	}
}