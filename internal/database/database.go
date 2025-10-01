package database

import (
	"context"
	"gobee/ent"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *ent.Client

func InitializeDB(autoMigrate bool) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	DB = client

	if autoMigrate {
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}
}

func CloseDB() {
	DB.Close()
}
