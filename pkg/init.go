package pkg

import "gobee/internal/database"

const autoMigrate = true

func InitializeServices() {
	database.InitializeDB(database.DBConfig{
		DBType:     "sqlite",
		SqliteFile: "",
	}, autoMigrate)
}
