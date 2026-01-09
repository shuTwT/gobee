package ftp

import (
	"io"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jlaffaye/ftp"
)

// FTPUploader 是FTP存储的上传器实现
type FTPUploader struct {
	host     string
	port     int
	username string
	password string
	basePath string
}

func (u *FTPUploader) formatAddress() string {
	if u.port == 21 {
		return u.host
	}
	return u.host + ":" + strconv.Itoa(u.port)
}

// NewFTPStorage 创建一个新的FTP存储实例
func NewFTPStorage(host string, port int, username string, password string, basePath string) (*FTPUploader, error) {
	if port == 0 {
		port = 21 // 默认FTP端口
	}

	return &FTPUploader{
		host:     host,
		port:     port,
		username: username,
		password: password,
		basePath: basePath,
	}, nil
}

// 建立FTP连接
func (u *FTPUploader) connect() (*ftp.ServerConn, error) {
	addr := u.formatAddress()
	conn, err := ftp.Dial(addr)
	if err != nil {
		return nil, err
	}

	// 登录
	if err := conn.Login(u.username, u.password); err != nil {
		conn.Quit() // 确保连接关闭
		return nil, err
	}

	// 如果有基础路径，切换到该路径
	if u.basePath != "" {
		if err := conn.ChangeDir(u.basePath); err != nil {
			// 尝试创建目录（如果不存在）
			if err := u.createRemoteDir(conn, u.basePath); err != nil {
				conn.Quit()
				return nil, err
			}
			// 再次尝试切换目录
			if err := conn.ChangeDir(u.basePath); err != nil {
				conn.Quit()
				return nil, err
			}
		}
	}

	return conn, nil
}

// 创建远程目录（支持多级目录）
func (u *FTPUploader) createRemoteDir(conn *ftp.ServerConn, dir string) error {
	// 分割路径为多个部分
	dirs := filepath.SplitList(dir)
	currentDir := ""

	for _, d := range dirs {
		if d == "" {
			continue
		}

		// 构建当前要创建的目录路径
		nextDir := filepath.Join(currentDir, d)

		// 尝试切换到目录，如果失败则创建
		if err := conn.ChangeDir(nextDir); err != nil {
			// 目录不存在，创建它
			if err := conn.MakeDir(d); err != nil {
				return err
			}
			// 切换到新创建的目录
			if err := conn.ChangeDir(d); err != nil {
				return err
			}
		}

		currentDir = nextDir
	}

	return nil
}

// Upload 上传文件到FTP服务器
func (u *FTPUploader) Upload(key string, reader io.Reader, size int64, contentType string) (string, error) {
	conn, err := u.connect()
	if err != nil {
		return "", err
	}
	defer conn.Quit()

	// 创建文件所在的目录
	keyDir := filepath.Dir(key)
	if keyDir != "." && keyDir != "" {
		if err := u.createRemoteDir(conn, keyDir); err != nil {
			return "", err
		}
	}

	// 上传文件
	if err := conn.Stor(key, reader); err != nil {
		return "", err
	}

	return key, nil
}

// GetURL 获取文件访问URL
func (u *FTPUploader) GetURL(key string, expires time.Duration) (string, error) {
	// FTP通常不直接提供HTTP访问URL，需要配置Web服务器映射FTP目录
	// 这里假设已经配置了对应的HTTP服务
	return u.host + ":" + strconv.Itoa(u.port) + "/" + key, nil
}

// Delete 从FTP服务器删除文件
func (u *FTPUploader) Delete(key string) error {
	conn, err := u.connect()
	if err != nil {
		return err
	}
	defer conn.Quit()

	return conn.Delete(key)
}
