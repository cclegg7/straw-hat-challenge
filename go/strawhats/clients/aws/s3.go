package aws

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	bucketName = "straw-hat-challenge"
)

type S3Client struct {
	uploader *s3manager.Uploader
}

func NewS3Client() (*S3Client, error) {
	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		return nil, err
	}

	s3Uploader := s3manager.NewUploader(s3Session)
	return &S3Client{
		uploader: s3Uploader,
	}, nil
}

func (c *S3Client) Upload(file io.Reader, name string, contentType string) (string, error) {
	uploadResponse, err := c.uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(name),
		Body:        file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	return uploadResponse.Location, nil
}
