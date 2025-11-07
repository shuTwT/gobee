package router

import (
	album_handler "gobee/internal/handlers/album"
	albumphoto_handler "gobee/internal/handlers/album_photo"
	api_interface_handler "gobee/internal/handlers/api_interface"
	auth_handler "gobee/internal/handlers/auth"
	file_handler "gobee/internal/handlers/file"
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

			apiV1.Get("/api-interface/all-routes", api_interface_handler.GetAllRoutes)

			// 系统设置接口
			apiV1.Get("/settings", setting_handler.GetSettings)
			// 登录身份验证中间件
			apiV1.Use(middleware.Protected())

			fileApi := apiV1.Group("/file")
			{
				fileApi.Get("/list", file_handler.ListFile)
				fileApi.Get("/page", file_handler.ListFilePage)
				fileApi.Get("/query/:id", file_handler.QueryFile)
				fileApi.Delete("/delete/:id", file_handler.DeleteFile)
				fileApi.Post("/upload", file_handler.Upload)
			}
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
				storageStrategyApi.Get("/list-all", storagestrategy_handler.ListStorageStrategyAll)
				storageStrategyApi.Post("/create", storagestrategy_handler.CreateStorageStrategy)
				storageStrategyApi.Put("/update/:id", storagestrategy_handler.UpdateStorageStrategy)
				storageStrategyApi.Get("/query/:id", storagestrategy_handler.QueryStorageStrategy)
				storageStrategyApi.Delete("/delete/:id", storagestrategy_handler.DeleteStorageStrategy)
				storageStrategyApi.Put("/default/:id", storagestrategy_handler.SetDefaultStorageStrategy)
			}
			albumApi := apiV1.Group("/album")
			{
				albumApi.Get("/list", album_handler.ListAlbum)
				albumApi.Post("/create", album_handler.CreateAlbum)
				albumApi.Put("/update/:id", album_handler.UpdateAlbum)
				albumApi.Get("/query/:id", album_handler.QueryAlbum)
				albumApi.Delete("/delete/:id", album_handler.DeleteAlbum)
			}
			albumPhotoApi := apiV1.Group("/album-photo")
			{
				albumPhotoApi.Get("/list", albumphoto_handler.ListAlbumPhoto)
				albumPhotoApi.Post("/create", albumphoto_handler.CreateAlbumPhoto)
				albumPhotoApi.Put("/update/:id", albumphoto_handler.UpdateAlbumPhoto)
				albumPhotoApi.Get("/query/:id", albumphoto_handler.QueryAlbumPhoto)
				albumPhotoApi.Delete("/delete/:id", albumphoto_handler.DeleteAlbumPhoto)
			}

		}
	}
}
