package models

import (
	"regexp"

	"github.com/braam76/auth-backend/api/v1/utils"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Username string `gorm:"not null;unique;type:char(10)"`
	Password string `gorm:"not null"`
}

func (u User) ValidateUser() []string {
	validationSteps := []utils.ValidationStep{
		{Stmt: len(u.Username) == 10, Err: "Username length is not 10 chars"},
		{Stmt: regexp.MustCompile(`[0-9]+$`).MatchString(u.Username), Err: "Username should contain only numbers"},
	}

	return utils.Validate(validationSteps)
}
