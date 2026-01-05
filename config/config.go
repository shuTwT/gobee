package config

import (
	"log"
	"os"
	"path/filepath"

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

func Init() {
	viper.SetDefault(DATABASE_URL, "file:ent?mode=memory&cache=shared&_fk=1")
	viper.SetDefault(PORT, "8000")
	viper.SetDefault(STAGE, "dev")

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./data")
	viper.AddConfigPath("$HOME/.gobee")
	viper.AutomaticEnv()

	for _, key := range configKeys {
		if viper.Get(key) == nil {
			panic("config key not found: " + key)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件不存在，在工作空间/data创建默认配置文件
			createDefaultConfig()
		} else {
			// 配置文件存在但读取失败
			panic("fatal error config file: " + err.Error())
		}
	} else {
		log.Printf("已加载配置文件: %s", viper.ConfigFileUsed())
	}
}

func createDefaultConfig() {
	// 获取当前工作目录
	workDir, err := os.Getwd()
	if err != nil {
		println("警告: 无法获取当前工作目录:", err.Error())
		return
	}

	configDir := filepath.Join(workDir, "data")
	configPath := filepath.Join(configDir, "config.toml")

	// 确保目录存在
	if err := os.MkdirAll(configDir, 0755); err != nil {
		log.Panicln("警告: 无法创建配置目录:", err.Error())
		return
	}

	// 设置配置值
	viper.Set("database.url", viper.GetString(DATABASE_URL))
	viper.Set("server.port", viper.GetString(PORT))
	viper.Set("server.stage", viper.GetString(STAGE))

	// 在可执行文件同级目录创建配置文件
	if err := viper.WriteConfigAs(configPath); err != nil {
		// 如果写入失败，记录警告但不影响程序运行
		println("警告: 无法创建默认配置文件:", err.Error())
	} else {
		println("已创建默认配置文件:", configPath)
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
