package database

import (
	"context"
	"fmt"
	"gobee/ent"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var DB *ent.Client

type DBConfig struct {
	DBType string

	// SQLite
	SqliteFile string

	// MySQL
	MysqlHost     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
	MysqlDatabase string

	// PostgreSQL
	PgHost     string
	PgPort     string
	PgUser     string
	PgPassword string
	PgDatabase string
}

func InitializeDB(cfg DBConfig, autoMigrate bool) (*ent.Client, error) {
	var client *ent.Client
	var err error

	switch cfg.DBType {
	case "sqlite":
		// 如果没有提供文件路径，则使用默认路径
		if cfg.SqliteFile == "" {
			var homeDir string
			homeDir, err = os.UserHomeDir()
			if err != nil {
				log.Fatalf("failed getting user home directory: %v", err)
			}
			dataDir := filepath.Join(homeDir, "data")
			if err = os.MkdirAll(dataDir, 0755); err != nil {
				log.Fatalf("创建 ~/data 目录失败: %v", err)
			}
			cfg.SqliteFile = filepath.Join(dataDir, "sql.db")
		}
		client, err = ent.Open("sqlite3", fmt.Sprintf("file:%s?cache=shared&_fk=1", cfg.SqliteFile))
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
			cfg.MysqlUser, cfg.MysqlPassword, cfg.MysqlHost, cfg.MysqlPort, cfg.MysqlDatabase)
		client, err = ent.Open("mysql", dsn)
	case "postgresql":
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPassword, cfg.PgDatabase)
		client, err = ent.Open("postgres", dsn)
	default:
		log.Fatalf("unsupported database type: %s", cfg.DBType)
	}

	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", cfg.DBType, err)
	}
	DB = client

	if autoMigrate {
		log.Println("Auto migrating schema...")
		if err = client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}
	return client, err
}

func CloseDB() {
	DB.Close()
}
