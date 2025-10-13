package local

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"
)

// LocalStorage 本地存储实现
type LocalStorage struct {
	RootDir string // 根目录
	BaseURL string // 基础URL，用于生成访问链接
}

// NewLocalStorage 创建本地存储实例
func NewLocalStorage(rootDir, baseURL string) (*LocalStorage, error) {
	// 确保根目录存在
	if err := os.MkdirAll(rootDir, 0755); err != nil {
		return nil, err
	}

	return &LocalStorage{
		RootDir: rootDir,
		BaseURL: baseURL,
	}, nil
}

// Upload 上传文件到本地存储
func (l *LocalStorage) Upload(path string, reader io.Reader, size int64, contentType string) (string, error) {
	fullPath := filepath.Join(l.RootDir, path)

	// 创建目录
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	// 创建文件
	file, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 写入内容
	written, err := io.Copy(file, reader)
	if err != nil {
		return "", err
	}

	if written != size {
		return "", errors.New("写入的字节数与预期不符")
	}

	return path, nil
}

// GetURL 获取本地文件访问URL
func (l *LocalStorage) GetURL(path string, expiry time.Duration) (string, error) {
	// 本地文件URL不受有效期影响
	return l.BaseURL + "/" + path, nil
}

// Delete 删除本地文件
func (l *LocalStorage) Delete(path string) error {
	fullPath := filepath.Join(l.RootDir, path)
	return os.Remove(fullPath)
}
