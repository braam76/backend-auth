package main

import (
	"github.com/braam76/auth-backend/api"
	"github.com/braam76/auth-backend/api/v1/database"
	"github.com/braam76/auth-backend/api/v1/database/models"
	"github.com/braam76/auth-backend/api/v1/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	database.InitDB()
	database.DB.AutoMigrate(
		&models.UserModel{},
	)

	app.Route("/api", func(router fiber.Router) {
		utils.InitSessionStore()
		app.Use(cors.New(cors.Config{
			AllowOrigins:     "http://localhost:4321",
			AllowCredentials: true,
		}))
		// In future, add v2 same as v1 if needed
		router.Route("/v1", api.V1)
	})

	app.Listen(":3000")
}
