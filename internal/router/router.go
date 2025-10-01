package router

import (
	auth_handlers "gobee/internal/handlers/auth"
	"gobee/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initialize(router *fiber.App) {
	router.Use(middleware.Security)
	router.Use(middleware.Authenticated)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	console := router.Group("/console")
	{
		console.Get("/", func(c *fiber.Ctx) error {
			return c.Render("pages/console/index", fiber.Map{
				"Title": "Hello, World!",
			}, "layouts/admin", "layouts/base")
		})
		console.Get("/login", func(c *fiber.Ctx) error {
			return c.Render("pages/console/login", fiber.Map{
				"Title": "登录",
			}, "layouts/base")
		})
	}

	auth := router.Group("/auth")
	{
		auth.Post("/login/password", auth_handlers.LoginPassword)
	}

	router.Get("/initialize", func(c *fiber.Ctx) error {
		return c.Render("pages/initialize", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/base")
	})
}
