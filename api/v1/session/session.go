package session

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var Store *sessions.CookieStore

func InitSessionStore() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// If needed, change to other store (db-based, cache-based, or smth)
	Store = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
}
