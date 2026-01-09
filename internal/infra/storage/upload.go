package storage

import (
	"io"
	"time"
)

// Uploader 定义上传接口，所有存储策略都需要实现此接口
type Uploader interface {
	// Upload 上传文件
	// path: 存储路径
	// reader: 数据读取器
	// size: 数据大小
	// contentType: 内容类型
	Upload(path string, reader io.Reader, size int64, contentType string) (string, error)

	// GetURL 获取文件访问URL
	// path: 存储路径
	// expiry: URL有效期，0表示永久有效
	GetURL(path string, expiry time.Duration) (string, error)

	// Delete 删除文件
	Delete(path string) error
}
