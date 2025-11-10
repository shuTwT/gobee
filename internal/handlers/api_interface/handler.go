package apiinterface

import (
	apiInterfaceService "gobee/internal/services/api_interface"

	"github.com/gofiber/fiber/v2"
)

func ListApiRoutesPage(c *fiber.Ctx) error {
	return apiInterfaceService.ListApiRoutesPage(c)
}
