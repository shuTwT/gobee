package storage

import (
	"context"
	"fmt"
	"gobee/ent"
	"gobee/ent/storagestrategy"
	"gobee/internal/database"
	"gobee/internal/infra/storage/ftp"
	"gobee/internal/infra/storage/local"
	"gobee/internal/infra/storage/s3"
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

var uploaderCache = map[int]Uploader{}

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

func ClearCache() {
	for id := range uploaderCache {
		delete(uploaderCache, id)
	}
}

// 获取上传器，如果不存在就新建
func GetUploader(strategy *ent.StorageStrategy) (Uploader, error) {
	if uploader, ok := uploaderCache[strategy.ID]; ok {
		return uploader, nil
	}

	// 从数据库获取配置
	storageStrategy, err := GetStorageStrategyByID(strategy.ID)
	if err != nil {
		return nil, err
	}

	var config StorageConfig

	if storageStrategy.Type == "s3" {
		config = StorageConfig{
			Type:       StorageS3,
			Endpoint:   storageStrategy.Endpoint,
			Region:     storageStrategy.Region,
			AccessKey:  storageStrategy.AccessKey,
			SecretKey:  storageStrategy.SecretKey,
			BucketName: storageStrategy.Bucket,
		}
	} else {
		config = StorageConfig{
			Type:     StorageLocal,
			BasePath: storageStrategy.BasePath,
			BaseURL:  "",
		}
	}

	// 创建上传器
	uploader, err := NewUploader(config)
	if err != nil {
		return nil, err
	}

	// 缓存上传器
	uploaderCache[strategy.ID] = uploader

	return uploader, nil
}

/**
 * GetMasterStorageStrategy 获取主存储策略
 * @return (*ent.StorageStrategy, error)
 */
func GetMasterStorageStrategy() (*ent.StorageStrategy, error) {
	client := database.DB
	// 找到 master 为 true 的
	strategy, err := client.StorageStrategy.Query().Where(storagestrategy.Master(true)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return strategy, nil
}

func GetStorageStrategyByID(id int) (*ent.StorageStrategy, error) {
	client := database.DB
	// 找到 master 为 true 的
	strategy, err := client.StorageStrategy.Query().Where(storagestrategy.ID(id)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return strategy, nil
}
