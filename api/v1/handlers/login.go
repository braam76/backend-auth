package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/braam76/auth-backend/api/v1/database/models"
	"github.com/braam76/auth-backend/api/v1/database/mysql"
	"github.com/braam76/auth-backend/api/v1/middleware"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error while decoding json: %s", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var dbUser models.User
	result := mysql.DB.Where("username = ?", user.Username).Find(&dbUser)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		log.Printf("Database error: %s", result.Error)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if user.Password != dbUser.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	session := r.Context().Value(middleware.ContextKey("session")).(*sessions.Session)
	session.Values["user_id"] = dbUser.ID
	session.Values["username"] = dbUser.Username
	
	if err := sessions.Save(r, w); err != nil {
		log.Printf("Error saving session: %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("%+v", session)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "Login successful"}
	json.NewEncoder(w).Encode(response)
}
