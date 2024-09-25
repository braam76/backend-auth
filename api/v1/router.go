package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/braam76/auth-backend/api/v1/database/mysql"
	"github.com/braam76/auth-backend/api/v1/models"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	err := mysql.InitDB()
	if err != nil {
		log.Fatalf("Error while init DB: %s", err)
	}

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "everything %s", "works")
	})
	r.Post("/create-user", createUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Error while decoding json: %s", err)
	}

	mysql.DB.Create(&user)
}


