package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	database = os.Getenv("MYSQL_DATABASE")
	user     = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	addr     = os.Getenv("MYSQL_ADDR")
	// rootPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
)

var DB *sql.DB
// TODO: rewrite with sqlx
func InitDB() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	

	if DB, err = sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, addr, database),
	); err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
}
