package storage

import (
	"fmt"
	"gobee/internal/services/storage/ftp"
	"gobee/internal/services/storage/local"
	"gobee/internal/services/storage/s3"
)

// StorageType 存储类型枚举
type StorageType string

const (
	StorageLocal   StorageType = "local"
	StorageFTP     StorageType = "ftp"
	StorageS3      StorageType = "s3"
	StorageAliyun  StorageType = "aliyun"
	StorageTencent StorageType = "tencent"
	StorageQiniu   StorageType = "qiniu"
	StorageHuawei  StorageType = "huawei"
)

// StorageConfig 存储配置
type StorageConfig struct {
	Type StorageType
	// 通用配置
	BasePath string
	BaseURL  string

	// FTP配置
	FTPHost     string
	FTPPort     int
	FTPUsername string
	FTPPassword string

	// 云存储通用配置
	Endpoint   string
	AccessKey  string
	SecretKey  string
	BucketName string
	Region     string
	Domain     string // 用于七牛云等需要自定义域名的存储
}

// NewUploader 根据配置创建对应的上传器
func NewUploader(config StorageConfig) (Uploader, error) {
	switch config.Type {
	case StorageLocal:
		return local.NewLocalStorage(config.BasePath, config.BaseURL)
	case StorageFTP:
		return ftp.NewFTPStorage(config.FTPHost, config.FTPPort, config.FTPUsername, config.FTPPassword, config.BasePath)
	case StorageS3:
		return s3.NewS3Storage(config.Endpoint, config.Region, config.AccessKey, config.SecretKey, config.BucketName)
	default:
		return nil, fmt.Errorf("不支持的存储类型: %s", config.Type)
	}
}
