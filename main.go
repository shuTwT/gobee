package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/mattn/go-sqlite3"

	"gobee/config"
	"gobee/internal/database"
	"gobee/internal/router"
	"gobee/pkg"

	"github.com/joho/godotenv"
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

	router.Initialize(app)
	return app
}

func main() {
	app := InitializeApp()
	defer database.CloseDB()
	app.Listen(":3000")
}
