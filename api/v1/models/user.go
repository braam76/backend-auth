package models

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Username string `gorm:"not null;unique;type:char(10)"`
	Password string	`gorm:"not null"`
}