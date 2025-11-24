package router

import (
	album_handler "gobee/internal/handlers/album"
	albumphoto_handler "gobee/internal/handlers/album_photo"
	api_interface_handler "gobee/internal/handlers/api_interface"
	auth_handler "gobee/internal/handlers/auth"
	comment_handler "gobee/internal/handlers/comment"
	file_handler "gobee/internal/handlers/file"
	initialize_handler "gobee/internal/handlers/initialize"
	paychannel_handler "gobee/internal/handlers/pay_channel"
	payorder_handler "gobee/internal/handlers/pay_order"
	post_handler "gobee/internal/handlers/post"
	role_handler "gobee/internal/handlers/role"
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

			apiV1.Get("/api-interface/page", api_interface_handler.ListApiRoutesPage)

			// 系统设置接口
			apiV1.Get("/settings", setting_handler.GetSettings)
			// 登录身份验证中间件
			apiV1.Use(middleware.Protected())

			settingsApi := apiV1.Group("/settings")
			{
				settingsApi.Get("/json/:key", setting_handler.GetJsonSettingsMap).Name("settingsMap")
				settingsApi.Post("/json/save/:key", setting_handler.SaveSettings).Name("jsonSettingsSave")
			}

			fileApi := apiV1.Group("/file")
			{
				fileApi.Get("/list", file_handler.ListFile).Name("fileList")
				fileApi.Get("/page", file_handler.ListFilePage).Name("filePage")
				fileApi.Get("/query/:id", file_handler.QueryFile).Name("fileQuery")
				fileApi.Delete("/delete/:id", file_handler.DeleteFile).Name("fileDelete")
				fileApi.Post("/upload", file_handler.Upload).Name("fileUpload")
			}
			payChannelApi := apiV1.Group("/pay-channel")
			{
				payChannelApi.Get("/list", paychannel_handler.ListPayChannel).Name("payChannelList")
				payChannelApi.Post("/create", paychannel_handler.CreatePayChannel).Name("payChannelCreate")
				payChannelApi.Put("/update/:id", paychannel_handler.UpdatePayChannel).Name("payChannelUpdate")
				payChannelApi.Get("/query/:id", paychannel_handler.QueryPayChannel).Name("payChannelQuery")
				payChannelApi.Delete("/delete/:id", paychannel_handler.DeletePayChannel).Name("payChannelDelete")
			}
			payOrderApi := apiV1.Group("/pay-order")
			{
				payOrderApi.Get("/list", payorder_handler.ListPayOrder).Name("payOrderList")
				payOrderApi.Put("/create", payorder_handler.CreatePayOrder).Name("payOrderCreate")
				payOrderApi.Put("/update/:id", payorder_handler.UpdatePayOrder).Name("payOrderUpdate")
				payOrderApi.Get("/query/:id", payorder_handler.QueryPayOrder).Name("payOrderQuery")
				payOrderApi.Delete("/delete/:id", payorder_handler.DeletePayOrder).Name("payOrderDelete")
			}
			roleApi := apiV1.Group("/role")
			{
				roleApi.Get("/list", role_handler.ListRole).Name("roleList")
				roleApi.Get("/page", role_handler.ListRolePage).Name("rolePage")
				roleApi.Post("/create", role_handler.CreateRole).Name("roleCreate")
				roleApi.Put("/update/:id", role_handler.UpdateRole).Name("roleUpdate")
				roleApi.Get("/query/:id", role_handler.QueryRole).Name("roleQuery")
				roleApi.Delete("/delete/:id", role_handler.DeleteRole).Name("roleDelete")
			}
			userApi := apiV1.Group("/user")
			{
				userApi.Get("/list", user_handler.ListUser).Name("userList")
				userApi.Get("/page", user_handler.ListUserPage).Name("userPage")
				userApi.Post("/create", user_handler.CreateUser).Name("userCreate")
				userApi.Put("/update/:id", user_handler.UpdateUser).Name("userUpdate")
				userApi.Get("/query/:id", user_handler.QueryUser).Name("userQuery")
				userApi.Delete("/delete/:id", user_handler.DeleteUser).Name("userDelete")
			}
			postApi := apiV1.Group("/post")
			{
				postApi.Get("/list", post_handler.ListPost).Name("postList")
				postApi.Post("/create", post_handler.CreatePost).Name("postCreate")
				postApi.Put("/update/content/:id", post_handler.UpdatePostContent).Name("postUpdateContent")
				postApi.Put("/update/setting/:id", post_handler.UpdatePostSetting).Name("postUpdateSetting")
				postApi.Put("/publish/:id", post_handler.PublishPost).Name("postPublish")
				postApi.Put("/unpublish/:id", post_handler.UnpublishPost).Name("postUnpublish")
				postApi.Get("/query/:id", post_handler.QueryPost).Name("postQuery")
				postApi.Delete("/delete/:id", post_handler.DeletePost).Name("postDelete")
			}
			commentApi := apiV1.Group("/comment")
			{
				commentApi.Get("/page", comment_handler.ListCommentPage).Name("commentPage")
			}
			storageStrategyApi := apiV1.Group("/storage-strategy")
			{
				storageStrategyApi.Get("/list", storagestrategy_handler.ListStorageStrategy).Name("storageStrategyList")
				storageStrategyApi.Get("/list-all", storagestrategy_handler.ListStorageStrategyAll).Name("storageStrategyListAll")
				storageStrategyApi.Post("/create", storagestrategy_handler.CreateStorageStrategy).Name("storageStrategyCreate")
				storageStrategyApi.Put("/update/:id", storagestrategy_handler.UpdateStorageStrategy).Name("storageStrategyUpdate")
				storageStrategyApi.Get("/query/:id", storagestrategy_handler.QueryStorageStrategy).Name("storageStrategyQuery")
				storageStrategyApi.Delete("/delete/:id", storagestrategy_handler.DeleteStorageStrategy).Name("storageStrategyDelete")
				storageStrategyApi.Put("/default/:id", storagestrategy_handler.SetDefaultStorageStrategy).Name("storageStrategyDefault")
			}
			albumApi := apiV1.Group("/album")
			{
				albumApi.Get("/list", album_handler.ListAlbum).Name("albumList")
				albumApi.Get("/page", album_handler.ListAlbumPage).Name("albumPage")
				albumApi.Post("/create", album_handler.CreateAlbum).Name("albumCreate")
				albumApi.Put("/update/:id", album_handler.UpdateAlbum).Name("albumUpdate")
				albumApi.Get("/query/:id", album_handler.QueryAlbum).Name("albumQuery")
				albumApi.Delete("/delete/:id", album_handler.DeleteAlbum).Name("albumDelete")
			}
			albumPhotoApi := apiV1.Group("/album-photo")
			{
				albumPhotoApi.Get("/list", albumphoto_handler.ListAlbumPhoto).Name("albumPhotoList")
				albumPhotoApi.Get("/page", albumphoto_handler.ListAlbumPhotoPage).Name("albumPhotoPage")
				albumPhotoApi.Post("/create", albumphoto_handler.CreateAlbumPhoto).Name("albumPhotoCreate")
				albumPhotoApi.Put("/update/:id", albumphoto_handler.UpdateAlbumPhoto).Name("albumPhotoUpdate")
				albumPhotoApi.Get("/query/:id", albumphoto_handler.QueryAlbumPhoto).Name("albumPhotoQuery")
				albumPhotoApi.Delete("/delete/:id", albumphoto_handler.DeleteAlbumPhoto).Name("albumPhotoDelete")
			}

		}
	}
}
