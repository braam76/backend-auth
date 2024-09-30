package dto

import (
	"regexp"

	"github.com/braam76/auth-backend/api/v1/utils"
)

type LoginUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u LoginUserDTO) Validate() []string {
	return utils.Validate([]utils.ValidationStep{
		{Stmt: len(u.Username) == 10, Error: "Username should be 10 chars long"},
		{Stmt: regexp.MustCompile(`^[0-9]+$`).MatchString(u.Username), Error: "Username should contain only numbers"},
	})
}
