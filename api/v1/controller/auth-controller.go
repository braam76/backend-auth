package controller

import (
	authservice "github.com/braam76/auth-backend/api/v1/service/auth-service"
	"github.com/gofiber/fiber/v2"
)

func AuthController(auth fiber.Router) {
	auth.Post("/login", authservice.Login)
	auth.Post("/create", authservice.Create)
}
