package models

import (
	"fmt"

	"github.com/braam76/auth-backend/api/v1/database"
)

var CreateUserTable = `
CREATE TABLE IF NOT EXISTS users (
	id 			INT AUTO_INCREMENT PRIMARY KEY,
	username	VARCHAR(50),
	password 	VARCHAR(100) UNIQUE
);`

func init() {
	if _, err := database.DB.Exec(CreateUserTable); err != nil {
		panic(err)
	}

	fmt.Println("Table 'users' created or already exists.")
}

func Insert(username, password string) {
	if _, err := database.DB.Query("INSERT INTO users VALUE('%s', '%s')")
}

func GetAll() {
	
}
