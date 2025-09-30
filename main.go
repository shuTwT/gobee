package main

import (
	"context"
	"gobee/ent"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine := html.New("./views", ".tmpl")
	engine.Debug(true)
	engine.Reload(true)
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	app := fiber.New(fiber.Config{
		AppName: "Fiber HTML Template Demo",
		Views:   engine, // 关联模板引擎
	})
	app.Static("/static", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/admin", "layouts/base")
	})

	app.Get("/initialize", func(c *fiber.Ctx) error {
		return c.Render("pages/initialize", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/base")
	})

	app.Listen(":3000")
}
