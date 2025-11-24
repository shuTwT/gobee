package main

import (
	"embed"

	_ "github.com/mattn/go-sqlite3"

	"gobee/cmd"
	"gobee/internal/database"

	_ "gobee/docs"
)

//go:embed assets/moduleDefs
var moduleDefs embed.FS

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
	app := cmd.InitializeApp(moduleDefs)
	defer database.CloseDB()
	app.Listen(":13000")
}
