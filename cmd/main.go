package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/braam76/auth-backend/api/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	r.Route("/api/v1", v1.Router)

	fmt.Print("Running on http://localhost:3000")
	log.Fatal("Error while running app", http.ListenAndServe(":3000", r))
}
