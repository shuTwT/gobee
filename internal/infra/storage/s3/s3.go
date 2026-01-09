package s3

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Storage S3兼容对象存储实现
type S3Storage struct {
	client     *s3.Client
	BucketName string
	Region     string
	Endpoint   string // 用于兼容S3的其他存储服务
}

// NewS3Storage 创建S3存储实例
func NewS3Storage(endpoint, region, accessKey, secretKey, bucketName string) (*S3Storage, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKey,
			secretKey,
			"",
		)),
	)
	if err != nil {
		return nil, err
	}

	var client *s3.Client
	if endpoint != "" {
		// 用于兼容S3的存储服务，如MinIO等
		client = s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(endpoint)
		})
	} else {
		client = s3.NewFromConfig(cfg)
	}

	return &S3Storage{
		client:     client,
		BucketName: bucketName,
		Region:     region,
		Endpoint:   endpoint,
	}, nil
}

// Upload 上传文件到S3存储
func (s *S3Storage) Upload(path string, reader io.Reader, size int64, contentType string) (string, error) {
	_, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(s.BucketName),
		Key:           aws.String(path),
		Body:          reader,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(contentType),
	})

	return path, err
}

// GetURL 获取S3文件访问URL
func (s *S3Storage) GetURL(path string, expiry time.Duration) (string, error) {
	// 如果expiry为0，则返回公共URL
	if expiry == 0 {
		return "https://" + s.BucketName + "." + s.Endpoint + "/" + path, nil
	}

	return "https://" + s.BucketName + "." + s.Endpoint + "/" + path, nil
}

// Delete 删除S3上的文件
func (s *S3Storage) Delete(path string) error {
	_, err := s.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(path),
	})
	return err
}
