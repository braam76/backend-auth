package main

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/braam76/auth-backend/api/v1"
	"github.com/braam76/auth-backend/api/v1/database/mysql"
	"github.com/braam76/auth-backend/api/v1/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	mysql.InitDB()
	err := mysql.DB.AutoMigrate(models.User{})
	if err != nil {
		log.Panicf("Error while migrating table(s): %s", err)
	}

	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	r.Route("/api/v1", v1.Router)

	fmt.Print("Running on http://localhost:3000")
	log.Panicf("Error while running app: %s\n", http.ListenAndServe(":3000", r))
}
