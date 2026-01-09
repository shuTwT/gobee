package router

import (
	"gobee/internal/handlers"
	"gobee/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func initSystemRouter(router fiber.Router, handlerMap handlers.HandlerMap) {
	settingsApi := router.Group("/settings")
	{
		settingsApi.Get("/json/:key", handlerMap.SettingHandler.GetJsonSettingsMap).Name("settingsMap")
		settingsApi.Post("/json/save/:key", handlerMap.SettingHandler.SaveSettings).Name("jsonSettingsSave")
	}
	roleApi := router.Group("/role")
	{
		roleApi.Get("/list", handlerMap.RoleHandler.ListRole).Name("roleList")
		roleApi.Get("/page", handlerMap.RoleHandler.ListRolePage).Name("rolePage")
		roleApi.Post("/create", handlerMap.RoleHandler.CreateRole).Name("roleCreate")
		roleApi.Put("/update/:id", handlerMap.RoleHandler.UpdateRole).Name("roleUpdate")
		roleApi.Get("/query/:id", handlerMap.RoleHandler.QueryRole).Name("roleQuery")
		roleApi.Delete("/delete/:id", handlerMap.RoleHandler.DeleteRole).Name("roleDelete")
	}
	userApi := router.Group("/user")
	{
		userApi.Get("/profile", handlerMap.UserHandler.GetUserProfile).Name("userProfile")
		userApi.Get("/list", handlerMap.UserHandler.ListUser).Name("userList")
		userApi.Get("/page", handlerMap.UserHandler.ListUserPage).Name("userPage")
		userApi.Post("/create", handlerMap.UserHandler.CreateUser).Name("userCreate")
		userApi.Put("/update/:id", handlerMap.UserHandler.UpdateUser).Name("userUpdate")
		userApi.Get("/query/:id", handlerMap.UserHandler.QueryUser).Name("userQuery")
		userApi.Delete("/delete/:id", handlerMap.UserHandler.DeleteUser).Name("userDelete")
	}
}

func initContentRouter(router fiber.Router, handlerMap handlers.HandlerMap) {
	commentApi := router.Group("/comment")
	{
		commentApi.Get("/page", handlerMap.CommentHandler.ListCommentPage).Name("commentPage")
		commentApi.Get("/query/:id", handlerMap.CommentHandler.GetComment).Name("commentQuery")
	}
	albumApi := router.Group("/album")
	{
		albumApi.Get("/list", handlerMap.AlbumHandler.ListAlbum).Name("albumList")
		albumApi.Get("/page", handlerMap.AlbumHandler.ListAlbumPage).Name("albumPage")
		albumApi.Post("/create", handlerMap.AlbumHandler.CreateAlbum).Name("albumCreate")
		albumApi.Put("/update/:id", handlerMap.AlbumHandler.UpdateAlbum).Name("albumUpdate")
		albumApi.Get("/query/:id", handlerMap.AlbumHandler.QueryAlbum).Name("albumQuery")
		albumApi.Delete("/delete/:id", handlerMap.AlbumHandler.DeleteAlbum).Name("albumDelete")
	}
	albumPhotoApi := router.Group("/album-photo")
	{
		albumPhotoApi.Get("/list", handlerMap.AlbumPhotoHandler.ListAlbumPhoto).Name("albumPhotoList")
		albumPhotoApi.Get("/page", handlerMap.AlbumPhotoHandler.ListAlbumPhotoPage).Name("albumPhotoPage")
		albumPhotoApi.Post("/create", handlerMap.AlbumPhotoHandler.CreateAlbumPhoto).Name("albumPhotoCreate")
		albumPhotoApi.Put("/update/:id", handlerMap.AlbumPhotoHandler.UpdateAlbumPhoto).Name("albumPhotoUpdate")
		albumPhotoApi.Get("/query/:id", handlerMap.AlbumPhotoHandler.QueryAlbumPhoto).Name("albumPhotoQuery")
		albumPhotoApi.Delete("/delete/:id", handlerMap.AlbumPhotoHandler.DeleteAlbumPhoto).Name("albumPhotoDelete")
	}
	flinkApi := router.Group("/flink")
	{
		flinkApi.Get("/list", handlerMap.FlinkHandler.ListFlink).Name("flinkList")
		flinkApi.Get("/page", handlerMap.FlinkHandler.ListFlinkPage).Name("flinkPage")
		flinkApi.Post("/create", handlerMap.FlinkHandler.CreateFlink).Name("flinkCreate")
		flinkApi.Put("/update/:id", handlerMap.FlinkHandler.UpdateFlink).Name("flinkUpdate")
		flinkApi.Get("/query/:id", handlerMap.FlinkHandler.QueryFlink).Name("flinkQuery")
		flinkApi.Delete("/delete/:id", handlerMap.FlinkHandler.DeleteFlink).Name("flinkDelete")
	}
	flinkGroupApi := router.Group("/flink-group")
	{
		flinkGroupApi.Get("/list", handlerMap.FlinkGroupHandler.ListFLinkGroup).Name("flinkGroupList")
		flinkGroupApi.Post("/create", handlerMap.FlinkGroupHandler.CreateFlinkGroup).Name("flinkGroupCreate")
		flinkGroupApi.Put("/update/:id", handlerMap.FlinkGroupHandler.UpdateFlinkGroup).Name("flinkGroupUpdate")
		flinkGroupApi.Delete("/delete/:id", handlerMap.FlinkGroupHandler.DeleteFLinkGroup).Name("flinkGroupDelete")
	}
	friendCircleRecordApi := router.Group("/friend-circle")
	{
		friendCircleRecordApi.Get("/page", handlerMap.FriendCircleHandler.ListFriendCircleRecordPage).Name("friendCircleRecordPage")
	}
	essayApi := router.Group("/essay")
	{
		essayApi.Get("/page", handlerMap.EssayHandler.GetEssayPage).Name("essayPage")
		essayApi.Post("/create", handlerMap.EssayHandler.CreateEssay).Name("essayCreate")
		essayApi.Put("/update/:id", handlerMap.EssayHandler.UpdateEssay).Name("essayUpdate")
		essayApi.Delete("/delete/:id", handlerMap.EssayHandler.DeleteEssay).Name("essayDelete")
	}
	postApi := router.Group("/post")
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
}

func initInfraRouter(router fiber.Router, handlerMap handlers.HandlerMap) {
	storageStrategyApi := router.Group("/storage-strategy")
	{
		storageStrategyApi.Get("/list", handlerMap.StorageStrategyHandler.ListStorageStrategy).Name("storageStrategyList")
		storageStrategyApi.Get("/list-all", handlerMap.StorageStrategyHandler.ListStorageStrategyAll).Name("storageStrategyListAll")
		storageStrategyApi.Post("/create", handlerMap.StorageStrategyHandler.CreateStorageStrategy).Name("storageStrategyCreate")
		storageStrategyApi.Put("/update/:id", handlerMap.StorageStrategyHandler.UpdateStorageStrategy).Name("storageStrategyUpdate")
		storageStrategyApi.Get("/query/:id", handlerMap.StorageStrategyHandler.QueryStorageStrategy).Name("storageStrategyQuery")
		storageStrategyApi.Delete("/delete/:id", handlerMap.StorageStrategyHandler.DeleteStorageStrategy).Name("storageStrategyDelete")
		storageStrategyApi.Put("/default/:id", handlerMap.StorageStrategyHandler.SetDefaultStorageStrategy).Name("storageStrategyDefault")
	}
	fileApi := router.Group("/file")
	{
		fileApi.Get("/list", handlerMap.FileHandler.ListFile).Name("fileList")
		fileApi.Get("/page", handlerMap.FileHandler.ListFilePage).Name("filePage")
		fileApi.Get("/query/:id", handlerMap.FileHandler.QueryFile).Name("fileQuery")
		fileApi.Delete("/delete/:id", handlerMap.FileHandler.DeleteFile).Name("fileDelete")
		fileApi.Post("/upload", handlerMap.FileHandler.Upload).Name("fileUpload")
	}
}

func initMallRouter(router fiber.Router, handlerMap handlers.HandlerMap) {
	productApi := router.Group("/product")
	{
		productApi.Get("/list", handlerMap.ProductHandler.ListProducts).Name("productList")
		productApi.Get("/page", handlerMap.ProductHandler.ListProductsPage).Name("productPage")
		productApi.Post("/create", handlerMap.ProductHandler.CreateProduct).Name("productCreate")
		productApi.Put("/update/:id", handlerMap.ProductHandler.UpdateProduct).Name("productUpdate")
		productApi.Get("/query/:id", handlerMap.ProductHandler.QueryProduct).Name("productQuery")
		productApi.Delete("/delete/:id", handlerMap.ProductHandler.DeleteProduct).Name("productDelete")
		productApi.Put("/batch", handlerMap.ProductHandler.BatchUpdateProducts).Name("productBatchUpdate")
		productApi.Post("/batch/delete", handlerMap.ProductHandler.BatchDeleteProducts).Name("productBatchDelete")
	}
	memberApi := router.Group("/member")
	{
		memberApi.Get("/query/:user_id", handlerMap.MemberHandler.QueryMember).Name("memberQuery")
		memberApi.Get("/page", handlerMap.MemberHandler.QueryMemberPage).Name("memberPage")
		memberApi.Post("/create", handlerMap.MemberHandler.CreateMember).Name("memberCreate")
		memberApi.Put("/update/:id", handlerMap.MemberHandler.UpdateMember).Name("memberUpdate")
		memberApi.Delete("/delete/:id", handlerMap.MemberHandler.DeleteMember).Name("memberDelete")
	}
	memberLevelApi := router.Group("/member-level")
	{
		memberLevelApi.Get("/query/:id", handlerMap.MemberLevelHandler.QueryMemberLevel).Name("memberLevelQuery")
		memberLevelApi.Get("/list", handlerMap.MemberLevelHandler.QueryMemberLevelList).Name("memberLevelList")
		memberLevelApi.Get("/page", handlerMap.MemberLevelHandler.QueryMemberLevelPage).Name("memberLevelPage")
		memberLevelApi.Post("/create", handlerMap.MemberLevelHandler.CreateMemberLevel).Name("memberLevelCreate")
		memberLevelApi.Put("/update/:id", handlerMap.MemberLevelHandler.UpdateMemberLevel).Name("memberLevelUpdate")
		memberLevelApi.Delete("/delete/:id", handlerMap.MemberLevelHandler.DeleteMemberLevel).Name("memberLevelDelete")
	}
	payOrderApi := router.Group("/pay-order")
	{
		payOrderApi.Get("/page", handlerMap.PayOrderHandler.ListPayOrderPage).Name("payOrderPage")
		payOrderApi.Put("/create", handlerMap.PayOrderHandler.CreatePayOrder).Name("payOrderCreate")
		payOrderApi.Put("/update/:id", handlerMap.PayOrderHandler.UpdatePayOrder).Name("payOrderUpdate")
		payOrderApi.Get("/query/:id", handlerMap.PayOrderHandler.QueryPayOrder).Name("payOrderQuery")
		payOrderApi.Delete("/delete/:id", handlerMap.PayOrderHandler.DeletePayOrder).Name("payOrderDelete")
		payOrderApi.Post("/submit", handlerMap.PayOrderHandler.SubmitPayOrder).Name("payOrderSubmit")
	}

	walletApi := router.Group("/wallet")
	{
		walletApi.Get("/query/:user_id", handlerMap.WalletHandler.QueryWallet).Name("walletQuery")
		walletApi.Get("/page", handlerMap.WalletHandler.QueryWalletPage).Name("walletPage")
		walletApi.Put("/update/:id", handlerMap.WalletHandler.UpdateWallet).Name("walletUpdate")
	}
}

func Initialize(router *fiber.App, handlerMap handlers.HandlerMap) {
	router.Use(middleware.Security)
	router.Get("/api/preinit", handlerMap.InitializeHandler.PreInit)
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
			apiV1.Get("/visit", handlerMap.VisitHandler.HandleVisitor)
			apiV1.Get("/routes", handlerMap.RouteHandler.GetRoutes)

			apiV1.Get("/api-interface/page", handlerMap.ApiInterfaceHandler.ListApiRoutesPage)

			// 系统设置接口
			apiV1.Get("/settings", handlerMap.SettingHandler.GetSettings)
			apiV1.All("/twikoo", handlerMap.CommentHandler.HandleTwikoo).Name("twikoo")
			apiV1.Get("/comment/recent", handlerMap.CommentHandler.RecentComment).Name("recentComment")
			apiV1.Get("/flink/random", handlerMap.FlinkHandler.RandomFlink).Name("randomFlink")
			// 登录身份验证中间件
			// apiV1.Use(middleware.Protected())

			// 首页统计信息接口
			apiV1.Get("/common/statistic", handlerMap.CommonHandler.GetHomeStatistics).Name("homeStatistic")

			apiV1.Get("/user/personal-access-token/list", handlerMap.UserHandler.GetPersonalAccessTokenList).Name("patList")
			apiV1.Get("/user/personal-access-token/query/:id", handlerMap.UserHandler.GetPersonalAccessToken).Name("patInfo")
			apiV1.Post("/user/personal-access-token/create", handlerMap.UserHandler.CreatePat).Name("patCreate")

			initContentRouter(apiV1, handlerMap)
			initInfraRouter(apiV1, handlerMap)
			initMallRouter(apiV1, handlerMap)
			initSystemRouter(apiV1, handlerMap)
		}
	}
}
