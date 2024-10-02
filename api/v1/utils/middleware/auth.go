package middleware

import (
	"log"

	"github.com/braam76/auth-backend/api/v1/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Get current session
	session, err := utils.SessionStore.Get(c)
	if err != nil {
		log.Printf("[ERROR] = %s\n", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	// If there is no session, dont let him go further
	userID := session.Get("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "not allowed",
		})
	}

	// If there is session with user
	return c.Next()

}
