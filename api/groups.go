package api

import (
	"github.com/braam76/auth-backend/api/v1/controller"
	"github.com/gofiber/fiber/v2"
)

func V1(router fiber.Router) {
	router.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"everything": "works",
		})
	})

	router.Route("/auth", controller.AuthController)
}
