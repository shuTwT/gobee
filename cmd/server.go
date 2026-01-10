package cmd

import (
	"embed"
	"gobee/internal/handlers"
	"gobee/internal/router"
	"gobee/pkg"
	"gobee/pkg/config"
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	"github.com/gofiber/template/html/v2"
)

func InitializeApp(moduleDefs embed.FS, frontendRes embed.FS) *fiber.App {
	config.Init()
	serviceMap := pkg.InitializeServices(moduleDefs)

	viewsPkg, _ := fs.Sub(moduleDefs, "views")
	engine := html.NewFileSystem(http.FS(viewsPkg), ".tmpl")
	engine.Debug(true)
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		AppName:       "Fiber HTML Template Demo",
		Prefork:       true,    // 启用多进程（Prefork 模式）
		CaseSensitive: true,    // 路由大小写敏感
		StrictRouting: true,    // 严格匹配带 / 和不带 / 的路由
		ServerHeader:  "Fiber", // 返回响应头中的 Server 字段
		Views:         engine,  // 关联模板引擎
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

	handlerMap := handlers.InitHandler(serviceMap)

	router.Initialize(app, handlerMap)

	return app
}
