package config

import (
	"os"

	"github.com/spf13/viper"
)

var configKeys = []string{
	DATABASE_URL,
	PORT,
	STAGE,
}

const (
	DATABASE_URL = "DATABASE_URL"
	PORT         = "PORT"
	STAGE        = "STAGE"
)

func init() {
	viper.SetDefault(DATABASE_URL, "file:ent?mode=memory&cache=shared&_fk=1")
	viper.SetDefault(PORT, "8000")

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/gobee")
	viper.AddConfigPath("$HOME/.gobee")
	viper.AutomaticEnv()

	for _, key := range configKeys {
		if viper.Get(key) == nil {
			panic("config key not found: " + key)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		panic("fatal error config file: " + err.Error())
	}
}

func GetDatabaseUrl() string {
	return os.Getenv("DATABASE_URL")
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8000"
	}
	return port
}

func GetIsProduction() bool {
	return os.Getenv("STAGE") == "prod"
}

func GetSecret() string {
	return os.Getenv("SECRET")
}
