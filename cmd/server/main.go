package cmd

import (
	"context"
	"embed"
	"fmt"
	"log/slog"

	"github.com/shuTwT/gobee/ent"
	"github.com/shuTwT/gobee/internal/database"
	"github.com/shuTwT/gobee/internal/handlers"
	"github.com/shuTwT/gobee/internal/router"
	"github.com/shuTwT/gobee/internal/schedule"
	"github.com/shuTwT/gobee/internal/schedule/manager"
	"github.com/shuTwT/gobee/pkg"
	"github.com/shuTwT/gobee/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func InitializeApp(moduleDefs embed.FS, frontendRes embed.FS) (*fiber.App, *ent.Client) {
	config.Init()
	dbType := config.GetString(config.DATABASE_TYPE)
	dbConfig := database.DBConfig{
		DBType: dbType,
		DBUrl:  config.GetString(config.DATABASE_URL),
	}
	db, err := database.InitializeDB(dbConfig, true)
	if err != nil {
		panic(err)
	}

	scheduleManager := manager.NewScheduleManager()

	serviceMap := pkg.InitializeServices(moduleDefs, db, scheduleManager)

	if !fiber.IsChild() {
		// 主进程程初始化定时任务
		err := schedule.InitializeSchedule(db, scheduleManager, serviceMap.FriendCircleService)
		if err != nil {
			defer scheduleManager.Shutdown()
		}
	}

	app := fiber.New(fiber.Config{
		AppName:       "Fiber HTML Template Demo",
		Prefork:       true,    // 启用多进程（Prefork 模式）
		CaseSensitive: true,    // 路由大小写敏感
		StrictRouting: true,    // 严格匹配带 / 和不带 / 的路由
		ServerHeader:  "Fiber", // 返回响应头中的 Server 字段
	})

	if config.GetBool(config.SWAGGER_ENABLE) {
		app.Get("/swagger/*", swagger.HandlerDefault) // default
	}

	// app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
	// 	URL:         "http://example.com/doc.json",
	// 	DeepLinking: false,
	// 	// Expand ("list") or Collapse ("none") tag groups by default
	// 	DocExpansion: "none",
	// 	// Prefill OAuth ClientId on Authorize popup
	// 	OAuth: &swagger.OAuthConfig{
	// 		AppName:  "OAuth Provider",
	// 		ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
	// 	},
	// 	// Ability to change OAuth2 redirect uri location
	// 	OAuth2RedirectUrl: "http://localhost:3000/swagger/oauth2-redirect.html",
	// }))
	app.Use(logger.New())
	app.Use(cors.New())

	router.InitFrontendRes(app, frontendRes)

	handlerMap := handlers.InitHandler(serviceMap, db)
	router.Initialize(app, handlerMap, db)

	go func() {
		if fiber.IsChild() {
			return
		}
		if err := serviceMap.PluginService.AutoStartPlugins(context.Background()); err != nil {
			fmt.Printf("自动启动插件失败: %v\n", err)
		}
		slog.Info("自动启动插件成功")
	}()

	return app, db
}
