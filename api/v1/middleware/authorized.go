package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

func AuthorizedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value(ContextKey("session")).(*sessions.Session)
		fmt.Printf("%+v", session)
		if session.Values["user_id"] == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{"msg": "Unauthorized"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
