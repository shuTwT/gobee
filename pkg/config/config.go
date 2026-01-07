package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var configKeys = []string{
	DATABASE_URL,
	SERVER_PORT,
	SERVER_STAGE,
}

const (
	DATABASE_TYPE         = "database.type"
	DATABASE_URL          = "database.url"
	SERVER_PORT           = "server.port"
	SERVER_STAGE          = "server.stage"
	SERVER_DEBUG          = "server.debug"
	SWAGGER_ENABLE        = "swagger.enable"
	AUTH_TOKEN_SECRET     = "auth.token_secret"
	AUTH_PAT_SECRET       = "auth.pat_secret"
	PAY_EPAY_ENABLE       = "pay.epay_enable"
	PAY_EPAY_API_URL      = "pay.epay_api_url"
	PAY_EPAY_MERCHANT_ID  = "pay.epay_merchant_id"
	PAY_EPAY_MERCHANT_KEY = "pay.epay_merchant_key"
	PAY_EPAY_NOTIFY_URL   = "pay.epay_notify_url"
	PAY_EPAY_RETURN_URL   = "pay.epay_return_url"
)

func Init() {
	viper.SetDefault(DATABASE_TYPE, "sqlite")
	viper.SetDefault(DATABASE_URL, "file:./data/sql.db?cache=shared&_fk=1")
	viper.SetDefault(SERVER_PORT, "8000")
	viper.SetDefault(SERVER_STAGE, "dev")
	viper.SetDefault(SERVER_DEBUG, false)
	viper.SetDefault(SWAGGER_ENABLE, true)
	viper.SetDefault(AUTH_TOKEN_SECRET, "your-secret-key")
	viper.SetDefault(AUTH_PAT_SECRET, "your-pat-secret")
	// 易支付相关
	viper.SetDefault(PAY_EPAY_ENABLE, true)
	viper.SetDefault(PAY_EPAY_API_URL, "https://api.pay.com")
	viper.SetDefault(PAY_EPAY_MERCHANT_ID, "your-merchant-id")
	viper.SetDefault(PAY_EPAY_MERCHANT_KEY, "your-merchant-key")
	viper.SetDefault(PAY_EPAY_NOTIFY_URL, "https://api.shhsu.com/api/v1/payorder/notify")
	viper.SetDefault(PAY_EPAY_RETURN_URL, "https://api.shhsu.com/api/v1/payorder/return")

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

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}
