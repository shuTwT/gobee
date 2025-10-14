package auth_handler

import "github.com/gofiber/fiber/v2"

func LoginPassword(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "login password",
	})
}
