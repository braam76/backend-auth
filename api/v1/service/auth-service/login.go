package authservice

import (
	"log"

	"github.com/braam76/auth-backend/api/v1/database"
	"github.com/braam76/auth-backend/api/v1/database/models"
	"github.com/braam76/auth-backend/api/v1/dto"
	"github.com/braam76/auth-backend/api/v1/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var userDto dto.LoginUserDTO
	var userModel models.UserModel

	if err := c.BodyParser(&userDto); err != nil {
		log.Printf("[ERROR] = %s\n", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if errors := userDto.Validate(); errors != nil {
		log.Printf("[ERROR] = %+v\n", errors)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": errors})
	}

	result := database.DB.
		Where(&models.UserModel{
			Username: userDto.Username,
			Password: userDto.Password,
		}).
		First(&userModel)

	if result.Error != nil {
		log.Printf("[ERROR] = %s\n", result.Error)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	session, err := utils.SessionStore.Get(c)
	if err != nil {
		log.Printf("[ERROR] = %s\n", err)
		return nil
	}

	session.Set("user_id", userModel.ID)

	if err := session.Save(); err != nil {
		log.Printf("[ERROR] = %s\n", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	log.Printf("%+v\n", session.Get("user_id"))

	log.Println(userModel)
	return c.SendString("GOOD! Login")
}
