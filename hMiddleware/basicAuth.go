package helpers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func BasicAuth() fiber.Handler {
	username := os.Getenv("BASIC_AUTH_USERNAME")
	password := os.Getenv("BASIC_AUTH_PASSWORD")

	if username == "" || password == "" {
		panic("BASIC_AUTH_USERNAME and BASIC_AUTH_PASSWORD must be set")
	}

	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			username: password,
		},
	})
}
