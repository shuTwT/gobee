package router

import (
	auth_handler "gobee/internal/handlers/auth"
	initialize_handler "gobee/internal/handlers/initialize"
	paychannel_handler "gobee/internal/handlers/pay_channel"
	payorder_handler "gobee/internal/handlers/pay_order"
	route_handler "gobee/internal/handlers/route"
	setting_handler "gobee/internal/handlers/setting"
	user_handler "gobee/internal/handlers/user"
	"gobee/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initialize(router *fiber.App) {
	router.Use(middleware.Security)

	router.Post("/api/initialize", initialize_handler.Initialize)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	auth := router.Group("/api/auth")
	{
		auth.Post("/login/password", auth_handler.Login)
	}

	api := router.Group("/api")
	{
		apiV1 := api.Group("/v1")
		{
			apiV1.Get("/users", func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{
					"message": "Hello, World!",
				})
			})
			apiV1.Get("/routes", route_handler.GetRoutes)

			// 系统设置接口
			apiV1.Get("/settings", setting_handler.GetSettings)
			// 登录身份验证中间件
			apiV1.Use(middleware.Protected())
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
			userApi := apiV1.Group("/user")
			{
				userApi.Get("/list", user_handler.ListUser)
				userApi.Get("/create", user_handler.CreateUser)
				userApi.Get("/update", user_handler.UpdateUser)
				userApi.Get("/query", user_handler.QueryUser)
				userApi.Get("/delete", user_handler.DeleteUser)
			}
		}
	}
}
