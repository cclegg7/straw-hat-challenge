package aws

import (
	"fmt"
	"github.com/cclegg7/straw-hat-challenge/configs"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Client interface {
	Upload(file io.Reader, name string, contentType string) (string, error)
}

type s3ClientImpl struct {
	uploader   *s3manager.Uploader
	bucketName string
}

func NewS3Client(configs *configs.FileStorage) (S3Client, error) {
	if !configs.UseS3 {
		return &localStorageClient{}, nil
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(configs.S3Region),
	})
	if err != nil {
		return nil, err
	}

	s3Uploader := s3manager.NewUploader(s3Session)
	return &s3ClientImpl{
		uploader:   s3Uploader,
		bucketName: configs.S3Bucket,
	}, nil
}

func (c *s3ClientImpl) Upload(file io.Reader, name string, contentType string) (string, error) {
	uploadResponse, err := c.uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(c.bucketName),
		Key:         aws.String(name),
		Body:        file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	return uploadResponse.Location, nil
}
