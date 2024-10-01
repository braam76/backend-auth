package authservice

import (
	"log"

	"github.com/braam76/auth-backend/api/v1/database"
	"github.com/braam76/auth-backend/api/v1/database/models"
	"github.com/gofiber/fiber/v2"
)

func TestProtected(c *fiber.Ctx) error {
	var userModel *models.UserModel

	session, err := database.Redis.Get(c)
	if err != nil {
		log.Printf("[ERROR] = %s\n", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	result := database.DB.Where(&models.UserModel{
		ID: session.Get("user_id").(int),
	}).First(&userModel)

	if result.Error != nil {
		log.Printf("[ERROR] = %s\n", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"username": userModel.Username,
		"password": userModel.Password,
	})
}
