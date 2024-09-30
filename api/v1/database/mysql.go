package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var (
		database = os.Getenv("MYSQL_DATABASE")
		user     = os.Getenv("MYSQL_USER")
		password = os.Getenv("MYSQL_PASSWORD")
		addr     = os.Getenv("MYSQL_ADDR")

		// rootPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
	)

	databaseFullLink := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, addr, database)

	DB, err = gorm.Open(mysql.Open(databaseFullLink), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
