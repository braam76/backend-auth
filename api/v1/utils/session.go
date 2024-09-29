package utils

import "github.com/gofiber/fiber/v2/middleware/session"

var SessionStore *session.Store

func InitSessionStore() {
	// If needed, change store
	SessionStore = session.New()
}