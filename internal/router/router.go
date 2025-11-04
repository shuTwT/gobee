package router

import (
	auth_handler "gobee/internal/handlers/auth"
	initialize_handler "gobee/internal/handlers/initialize"
	paychannel_handler "gobee/internal/handlers/pay_channel"
	payorder_handler "gobee/internal/handlers/pay_order"
	post_handler "gobee/internal/handlers/post"
	route_handler "gobee/internal/handlers/route"
	setting_handler "gobee/internal/handlers/setting"
	storagestrategy_handler "gobee/internal/handlers/storage_strategy"
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
				payChannelApi.Post("/create", paychannel_handler.CreatePayChannel)
				payChannelApi.Put("/update/:id", paychannel_handler.UpdatePayChannel)
				payChannelApi.Get("/query/:id", paychannel_handler.QueryPayChannel)
				payChannelApi.Delete("/delete/:id", paychannel_handler.DeletePayChannel)
			}
			payOrderApi := apiV1.Group("/pay-order")
			{
				payOrderApi.Get("/list", payorder_handler.ListPayOrder)
				payOrderApi.Put("/create", payorder_handler.CreatePayOrder)
				payOrderApi.Put("/update/:id", payorder_handler.UpdatePayOrder)
				payOrderApi.Get("/query/:id", payorder_handler.QueryPayOrder)
				payOrderApi.Delete("/delete/:id", payorder_handler.DeletePayOrder)
			}
			userApi := apiV1.Group("/user")
			{
				userApi.Get("/list", user_handler.ListUser)
				userApi.Post("/create", user_handler.CreateUser)
				userApi.Put("/update/:id", user_handler.UpdateUser)
				userApi.Get("/query/:id", user_handler.QueryUser)
				userApi.Delete("/delete/:id", user_handler.DeleteUser)
			}
			postApi := apiV1.Group("/post")
			{
				postApi.Get("/list", post_handler.ListPost)
				postApi.Post("/create", post_handler.CreatePost)
				postApi.Put("/update/:id", post_handler.UpdatePost)
				postApi.Get("/query/:id", post_handler.QueryPost)
				postApi.Delete("/delete/:id", post_handler.DeletePost)
			}
			storageStrategyApi := apiV1.Group("/storage-strategy")
			{
				storageStrategyApi.Get("/list", storagestrategy_handler.ListStorageStrategy)
				storageStrategyApi.Post("/create", storagestrategy_handler.CreateStorageStrategy)
				storageStrategyApi.Put("/update/:id", storagestrategy_handler.UpdateStorageStrategy)
				storageStrategyApi.Get("/query/:id", storagestrategy_handler.QueryStorageStrategy)
				storageStrategyApi.Delete("/delete/:id", storagestrategy_handler.DeleteStorageStrategy)
				storageStrategyApi.Put("/default/:id", storagestrategy_handler.SetDefaultStorageStrategy)
			}
		}
	}
}
