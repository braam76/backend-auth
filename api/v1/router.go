package v1

import (
	"fmt"
	"net/http"

	"github.com/braam76/auth-backend/api/v1/handlers"
	"github.com/braam76/auth-backend/api/v1/middleware"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Post("/create-user", handlers.CreateUserHandler)
	r.Post("/login", handlers.LoginHandler)

	// only if authorized
	r.Route("/authorized", func(r chi.Router) {
		r.Use(middleware.AuthorizedMiddleware)
		r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "everything %s", "works")
		})
	})
}
