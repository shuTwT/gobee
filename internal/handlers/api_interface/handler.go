package apiinterface

import (
	"gobee/pkg/domain/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllRoutes(c *fiber.Ctx) error {
	app := *c.App()
	routes := app.GetRoutes()

	return c.JSON(model.NewSuccess("路由列表获取成功", routes))
}
