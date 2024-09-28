package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/braam76/auth-backend/api/v1/session"
)

type ContextKey string

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := session.Store.Get(r, "session_id")
		if err != nil {
			log.Printf("Error while getting session: %s", err)
			http.Error(w, "Could not get session", http.StatusInternalServerError)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, ContextKey("session"), session)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

		if err := session.Save(r, w); err != nil {
			log.Println("Error saving session:", err)
		}
	})
}
