package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/mattn/go-sqlite3"

	"gobee/config"
	"gobee/internal/database"
	"gobee/internal/router"
	"gobee/pkg"

	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	_ "gobee/docs"
)

func InitializeApp() *fiber.App {
	godotenv.Load()
	config.Init()
	pkg.InitializeServices()
	engine := html.New("./views", ".tmpl")
	engine.Debug(true)
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		AppName: "Fiber HTML Template Demo",
		Views:   engine, // 关联模板引擎
	})
	app.Static("/", "./assets")

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:3000/swagger/oauth2-redirect.html",
	}))

	router.Initialize(app)

	return app
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	app := InitializeApp()
	defer database.CloseDB()
	app.Listen(":3000")
}
