package route

import (
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
)

func GetRoutes(c *fiber.Ctx) error {
	return c.JSON(model.NewSuccess("success", []string{}))
}
