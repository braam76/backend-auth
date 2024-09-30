package main

import (
	"github.com/braam76/auth-backend/api"
	"github.com/braam76/auth-backend/api/v1/database"
	"github.com/braam76/auth-backend/api/v1/database/models"
	"github.com/braam76/auth-backend/api/v1/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.InitDB()
	database.DB.AutoMigrate(
		&models.UserModel{},
	)

	app.Route("/api", func(router fiber.Router) {
		utils.InitSessionStore()

		// In future, add v2 same as v1 if needed
		router.Route("/v1", api.V1)
	})

	app.Listen(":3000")
}
