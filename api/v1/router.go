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
	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "everything %s", "works")
	})
	r.Post("/create-user", createUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Panicf("Error while decoding json: %s", err)
	}

	result := mysql.DB.Where("username = ?", user.Username).Find(&user)
	fmt.Print(result)
	if result.Error != nil {
		log.Printf("Error: %s", result.Error)
		return
	}

	if result.RowsAffected != 0 {
		log.Printf("Duplication error: \"%s\" already exists", user.Username)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	result = mysql.DB.Create(&user)
	if result.Error != nil {
		log.Panicf("Error while creating a user: %s", result.Error)
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
