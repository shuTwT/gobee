package router

import (
	auth_handler "gobee/internal/handlers/auth"
	paychannel_handler "gobee/internal/handlers/pay_channel"
	payorder_handler "gobee/internal/handlers/pay_order"
	"gobee/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initialize(router *fiber.App) {
	router.Use(middleware.Security)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	api := router.Group("/api")
	{
		api.Use(middleware.Protected())
		apiV1 := api.Group("/v1")
		{
			apiV1.Get("/users", func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{
					"message": "Hello, World!",
				})
			})
			payChannelApi := apiV1.Group("/pay-channel")
			{
				payChannelApi.Get("/list", paychannel_handler.ListPayChannel)
				payChannelApi.Get("/create", paychannel_handler.CreatePayChannel)
				payChannelApi.Get("/update", paychannel_handler.UpdatePayChannel)
				payChannelApi.Get("/query", paychannel_handler.QueryPayChannel)
				payChannelApi.Get("/delete", paychannel_handler.DeletePayChannel)
			}
			payOrderApi := apiV1.Group("/pay-order")
			{
				payOrderApi.Get("/list", payorder_handler.ListPayOrder)
				payOrderApi.Get("/create", payorder_handler.CreatePayOrder)
				payOrderApi.Get("/update", payorder_handler.UpdatePayOrder)
				payOrderApi.Get("/query", payorder_handler.QueryPayOrder)
				payOrderApi.Get("/delete", payorder_handler.DeletePayOrder)
			}
		}
	}

	auth := router.Group("/auth")
	{
		auth.Post("/login/password", auth_handler.Login)
	}

	router.Get("/initialize", func(c *fiber.Ctx) error {
		return c.Render("pages/initialize", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/base")
	})
}
