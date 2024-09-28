package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/braam76/auth-backend/api/v1/database/mysql"
	"github.com/braam76/auth-backend/api/v1/database/models"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        log.Printf("Error while decoding json: %s", err)
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

	result := mysql.DB.Where("username = ?", user.Username).Find(&user)

	fmt.Printf("result: %+v", result)
	if result.Error != nil {
		log.Printf("Database error: %s", result.Error)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected != 0 {
		log.Printf("Duplication error: \"%s\" already exists", user.Username)
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	if errors := user.ValidateUser(); errors != nil {
		log.Printf("Bad request: %s", errors)
		http.Error(w, fmt.Sprintf("Bad request: %s", errors), http.StatusBadRequest)
		return
	}

	if result = mysql.DB.Create(&user); result.Error != nil {
		log.Printf("Error while creating a user: %s", result.Error)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		log.Printf("Error encoding response: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
