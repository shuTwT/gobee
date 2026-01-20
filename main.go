package main

import (
	"embed"
	"log"

	_ "github.com/mattn/go-sqlite3"

	server "gobee/cmd/server"

	_ "gobee/docs"
	"gobee/internal/infra/logger"
)

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
	logger.NewLogger()
	app := server.InitializeApp(moduleDefs, frontendRes)

	log.Fatal(app.Listen(":13000"))
}
