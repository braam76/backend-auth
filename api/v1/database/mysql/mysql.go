package mysql

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() (err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_ADDR"),
		os.Getenv("MYSQL_DATABASE"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	return
}
