package main

import (
	"embed"

	_ "github.com/mattn/go-sqlite3"

	"gobee/cmd"
	"gobee/internal/database"
	"gobee/internal/schedule"

	_ "gobee/docs"
)

//go:embed views
//go:embed assets/moduleDefs
var moduleDefs embed.FS

//go:embed ui/dist
var frontendRes embed.FS

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:13000
// @BasePath /
func main() {
	app := cmd.InitializeApp(moduleDefs, frontendRes)
	defer database.CloseDB()
	scheduler, err := schedule.InitializeSchedule()
	if err != nil {
		defer scheduler.Shutdown()
	}

	app.Listen(":13000")
}
