package router

import (
	"gobee/internal/handlers"
	storagestrategy_handler "gobee/internal/handlers/storage_strategy"
	"gobee/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initialize(router *fiber.App, handlerMap handlers.HandlerMap) {
	router.Use(middleware.Security)
	router.Post("/api/initialize", handlerMap.InitializeHandler.Initialize)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	auth := router.Group("/api/auth")
	{
		auth.Post("/login/password", handlerMap.AuthHandler.Login)
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
			apiV1.Get("/routes", handlerMap.RouteHandler.GetRoutes)

			apiV1.Get("/api-interface/page", handlerMap.ApiInterfaceHandler.ListApiRoutesPage)

			// 系统设置接口
			apiV1.Get("/settings", handlerMap.SettingHandler.GetSettings)
			apiV1.All("/twikoo", handlerMap.CommentHandler.HandleTwikoo).Name("twikoo")
			apiV1.Get("/comment/recent", handlerMap.CommentHandler.RecentComment).Name("recentComment")
			// 登录身份验证中间件
			// apiV1.Use(middleware.Protected())

			// 首页统计信息接口
			apiV1.Get("/statistic", handlerMap.CommonHandler.GetHomeStatistics).Name("homeStatistic")

			apiV1.Get("/user/personal-access-token", handlerMap.UserHandler.GetPersonalAccessTokenList).Name("patList")
			apiV1.Get("/user/personal-access-token/:id", handlerMap.UserHandler.GetPersonalAccessToken).Name("patInfo")
			apiV1.Post("/user/personal-access-token", handlerMap.UserHandler.CreatePat).Name("patCreate")

			settingsApi := apiV1.Group("/settings")
			{
				settingsApi.Get("/json/:key", handlerMap.SettingHandler.GetJsonSettingsMap).Name("settingsMap")
				settingsApi.Post("/json/save/:key", handlerMap.SettingHandler.SaveSettings).Name("jsonSettingsSave")
			}

			fileApi := apiV1.Group("/file")
			{
				fileApi.Get("/list", handlerMap.FileHandler.ListFile).Name("fileList")
				fileApi.Get("/page", handlerMap.FileHandler.ListFilePage).Name("filePage")
				fileApi.Get("/query/:id", handlerMap.FileHandler.QueryFile).Name("fileQuery")
				fileApi.Delete("/delete/:id", handlerMap.FileHandler.DeleteFile).Name("fileDelete")
				fileApi.Post("/upload", handlerMap.FileHandler.Upload).Name("fileUpload")
			}
			payChannelApi := apiV1.Group("/pay-channel")
			{
				payChannelApi.Get("/list", handlerMap.PayChannelHandler.ListPayChannel).Name("payChannelList")
				payChannelApi.Post("/create", handlerMap.PayChannelHandler.CreatePayChannel).Name("payChannelCreate")
				payChannelApi.Put("/update/:id", handlerMap.PayChannelHandler.UpdatePayChannel).Name("payChannelUpdate")
				payChannelApi.Get("/query/:id", handlerMap.PayChannelHandler.QueryPayChannel).Name("payChannelQuery")
				payChannelApi.Delete("/delete/:id", handlerMap.PayChannelHandler.DeletePayChannel).Name("payChannelDelete")
			}
			payOrderApi := apiV1.Group("/pay-order")
			{
				payOrderApi.Get("/list", handlerMap.PayOrderHandler.ListPayOrder).Name("payOrderList")
				payOrderApi.Put("/create", handlerMap.PayOrderHandler.CreatePayOrder).Name("payOrderCreate")
				payOrderApi.Put("/update/:id", handlerMap.PayOrderHandler.UpdatePayOrder).Name("payOrderUpdate")
				payOrderApi.Get("/query/:id", handlerMap.PayOrderHandler.QueryPayOrder).Name("payOrderQuery")
				payOrderApi.Delete("/delete/:id", handlerMap.PayOrderHandler.DeletePayOrder).Name("payOrderDelete")
			}
			roleApi := apiV1.Group("/role")
			{
				roleApi.Get("/list", handlerMap.RoleHandler.ListRole).Name("roleList")
				roleApi.Get("/page", handlerMap.RoleHandler.ListRolePage).Name("rolePage")
				roleApi.Post("/create", handlerMap.RoleHandler.CreateRole).Name("roleCreate")
				roleApi.Put("/update/:id", handlerMap.RoleHandler.UpdateRole).Name("roleUpdate")
				roleApi.Get("/query/:id", handlerMap.RoleHandler.QueryRole).Name("roleQuery")
				roleApi.Delete("/delete/:id", handlerMap.RoleHandler.DeleteRole).Name("roleDelete")
			}
			userApi := apiV1.Group("/user")
			{
				userApi.Get("/profile", handlerMap.UserHandler.GetUserProfile).Name("userProfile")
				userApi.Get("/list", handlerMap.UserHandler.ListUser).Name("userList")
				userApi.Get("/page", handlerMap.UserHandler.ListUserPage).Name("userPage")
				userApi.Post("/create", handlerMap.UserHandler.CreateUser).Name("userCreate")
				userApi.Put("/update/:id", handlerMap.UserHandler.UpdateUser).Name("userUpdate")
				userApi.Get("/query/:id", handlerMap.UserHandler.QueryUser).Name("userQuery")
				userApi.Delete("/delete/:id", handlerMap.UserHandler.DeleteUser).Name("userDelete")
			}
			postApi := apiV1.Group("/post")
			{
				postApi.Get("/list", handlerMap.PostHandler.ListPost).Name("postList")
				postApi.Post("/create", handlerMap.PostHandler.CreatePost).Name("postCreate")
				postApi.Put("/update/content/:id", handlerMap.PostHandler.UpdatePostContent).Name("postUpdateContent")
				postApi.Put("/update/setting/:id", handlerMap.PostHandler.UpdatePostSetting).Name("postUpdateSetting")
				postApi.Put("/publish/:id", handlerMap.PostHandler.PublishPost).Name("postPublish")
				postApi.Put("/unpublish/:id", handlerMap.PostHandler.UnpublishPost).Name("postUnpublish")
				postApi.Get("/query/:id", handlerMap.PostHandler.QueryPost).Name("postQuery")
				postApi.Delete("/delete/:id", handlerMap.PostHandler.DeletePost).Name("postDelete")
				postApi.Get("/summary/stream/:id", handlerMap.PostHandler.GetSummaryForStream).Name("postSummaryStream")
			}
			commentApi := apiV1.Group("/comment")
			{
				commentApi.Get("/page", handlerMap.CommentHandler.ListCommentPage).Name("commentPage")
				commentApi.Get("/query/:id", handlerMap.CommentHandler.GetComment).Name("commentQuery")
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
				albumApi.Get("/list", handlerMap.AlbumHandler.ListAlbum).Name("albumList")
				albumApi.Get("/page", handlerMap.AlbumHandler.ListAlbumPage).Name("albumPage")
				albumApi.Post("/create", handlerMap.AlbumHandler.CreateAlbum).Name("albumCreate")
				albumApi.Put("/update/:id", handlerMap.AlbumHandler.UpdateAlbum).Name("albumUpdate")
				albumApi.Get("/query/:id", handlerMap.AlbumHandler.QueryAlbum).Name("albumQuery")
				albumApi.Delete("/delete/:id", handlerMap.AlbumHandler.DeleteAlbum).Name("albumDelete")
			}
			albumPhotoApi := apiV1.Group("/album-photo")
			{
				albumPhotoApi.Get("/list", handlerMap.AlbumPhotoHandler.ListAlbumPhoto).Name("albumPhotoList")
				albumPhotoApi.Get("/page", handlerMap.AlbumPhotoHandler.ListAlbumPhotoPage).Name("albumPhotoPage")
				albumPhotoApi.Post("/create", handlerMap.AlbumPhotoHandler.CreateAlbumPhoto).Name("albumPhotoCreate")
				albumPhotoApi.Put("/update/:id", handlerMap.AlbumPhotoHandler.UpdateAlbumPhoto).Name("albumPhotoUpdate")
				albumPhotoApi.Get("/query/:id", handlerMap.AlbumPhotoHandler.QueryAlbumPhoto).Name("albumPhotoQuery")
				albumPhotoApi.Delete("/delete/:id", handlerMap.AlbumPhotoHandler.DeleteAlbumPhoto).Name("albumPhotoDelete")
			}
			flinkApi := apiV1.Group("/flink")
			{
				flinkApi.Get("/list", handlerMap.FlinkHandler.ListFlink).Name("flinkList")
				flinkApi.Get("/page", handlerMap.FlinkHandler.ListFlinkPage).Name("flinkPage")
				flinkApi.Post("/create", handlerMap.FlinkHandler.CreateFlink).Name("flinkCreate")
				flinkApi.Put("/update/:id", handlerMap.FlinkHandler.UpdateFlink).Name("flinkUpdate")
				flinkApi.Get("/query/:id", handlerMap.FlinkHandler.QueryFlink).Name("flinkQuery")
				flinkApi.Delete("/delete/:id", handlerMap.FlinkHandler.DeleteFlink).Name("flinkDelete")
			}
			flinkGroupApi := apiV1.Group("/flink-group")
			{
				flinkGroupApi.Get("/list", handlerMap.FlinkGroupHandler.ListFLinkGroup).Name("flinkGroupList")
				flinkGroupApi.Post("/create", handlerMap.FlinkGroupHandler.CreateFlinkGroup).Name("flinkGroupCreate")
				flinkGroupApi.Put("/update/:id", handlerMap.FlinkGroupHandler.UpdateFlinkGroup).Name("flinkGroupUpdate")
				flinkGroupApi.Delete("/delete/:id", handlerMap.FlinkGroupHandler.DeleteFLinkGroup).Name("flinkGroupDelete")
			}
			friendCircleRuleApi := apiV1.Group("/friend-circle-rule")
			{
				friendCircleRuleApi.Get("/list", handlerMap.FriendCircleRuleHandler.ListFriendCircleRule).Name("friendCiecleRuleList")
				friendCircleRuleApi.Get("/page", handlerMap.FriendCircleRuleHandler.ListFriendCircleRulePage).Name("friendCiecleRulePage")
				friendCircleRuleApi.Post("/create", handlerMap.FriendCircleRuleHandler.CreateFriendCircleRule).Name("friendCiecleRuleCreate")
				friendCircleRuleApi.Put("/update/:id", handlerMap.FriendCircleRuleHandler.UpdateFriendCircleRule).Name("friendCiecleRuleUpdate")
				friendCircleRuleApi.Delete("/delete/:id", handlerMap.FriendCircleRuleHandler.DeleteFriendCircleRule).Name("friendCiecleRuleDelete")
			}
			friendCircleRecordApi := apiV1.Group("/friend-circle-record")
			{
				friendCircleRecordApi.Get("/page", handlerMap.FriendCircleRecordHandler.ListFriendCircleRecordPage).Name("friendCircleRecordPage")
			}

		}
	}
}
