package authservice

import (
	"log"

	"github.com/braam76/auth-backend/api/v1/database"
	"github.com/braam76/auth-backend/api/v1/database/models"
	"github.com/braam76/auth-backend/api/v1/dto"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	var userDto dto.LoginUserDTO

	if err := c.BodyParser(&userDto); err != nil {
		log.Printf("[ERROR] = %s\n", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	result := database.DB.Create(
		&models.UserModel{
			Username: userDto.Username,
			Password: userDto.Password,
		},
	)

	if result.Error != nil {
		log.Printf("[ERROR] = %s\n", result.Error)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	log.Printf("%+v\n", result)

	return c.SendString("GOOD! Create")
}