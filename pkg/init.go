package pkg

import (
	"embed"
	"gobee/internal/database"
	permission_service "gobee/internal/services/permission"
)

const autoMigrate = true

func InitializeServices(moduleDefs embed.FS) {
	database.InitializeDB(database.DBConfig{
		DBType:     "sqlite",
		SqliteFile: "",
	}, autoMigrate)
	permission_service.LoadPermissionsFromDef(moduleDefs)
}
