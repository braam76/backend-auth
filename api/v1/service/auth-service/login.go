package authservice

import (
	"fmt"
	"log"

	"github.com/braam76/auth-backend/api/v1/dto"
	"github.com/braam76/auth-backend/api/v1/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	userDto := dto.LoginUserDTO{}
	if err := c.BodyParser(&userDto); err != nil {
		log.Printf("[ERROR] = %s", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	

	session, err := utils.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", userDto)
	return c.SendString("GOOD!")
}
