package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authenticated(c *fiber.Ctx) error {
	if c.Locals("user") != nil {
		return c.Next()
	}

	if !strings.HasPrefix(c.Path(), "/console") {
		return c.Next()
	}

	if c.Path() == "/console/login" {
		return c.Next()
	}

	// session := c.Cookies("session")

	return c.Redirect("/login")

}
