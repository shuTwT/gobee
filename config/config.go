package config

import (
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
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件不存在，在可执行文件同级目录创建默认配置文件
			createDefaultConfig()
		} else {
			// 配置文件存在但读取失败
			panic("fatal error config file: " + err.Error())
		}
	}
}

func createDefaultConfig() {
	// 获取可执行文件路径
	execPath, err := os.Executable()
	if err != nil {
		println("警告: 无法获取可执行文件路径:", err.Error())
		return
	}

	// 获取可执行文件所在目录
	execDir := filepath.Dir(execPath)
	configPath := filepath.Join(execDir, "config.toml")

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
