package controller

import (
	authservice "github.com/braam76/auth-backend/api/v1/service/auth-service"
	"github.com/braam76/auth-backend/api/v1/utils/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthController(auth fiber.Router) {
	auth.Post("/create", authservice.Create)
	auth.Post("/login", authservice.Login)
	auth.Use(middleware.AuthMiddleware)                    // Apply middleware here
	auth.Get("/test-protected", authservice.TestProtected) // This route requires authentication
}
