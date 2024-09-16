package s3client

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Client struct {
	client *s3.S3
}

// NewS3Client initializes a new S3 client using AWS SDK v1
func NewS3Client(cfg config.Config) *s3Client {
	// Get credentials from environment variables
	accessKey := cfg.GetAws().AccessKeyId
	secretKey := cfg.GetAws().SecretAccessKey
	region := cfg.GetAws().Region

	if accessKey == "" || secretKey == "" || region == "" {
		panic("AWS credentials or region not found in .env file")
	}

	// Create the AWS credentials object
	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")

	// Create a new session with AWS SDK v1
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(region),
		Credentials:      creds,
		Endpoint:         aws.String("https://storage.googleapis.com"),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to create session: %v", err))
	}

	// Create the S3 client
	client := s3.New(sess)

	return &s3Client{
		client: client,
	}
}

func (c *s3Client) UploadFile(bucketName, objectKey string, buffer *bytes.Reader) error {
	_, err := c.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   buffer,
		ACL:    aws.String(s3.ObjectCannedACLPublicRead),
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}

	return nil
}

func (c *s3Client) DownloadFile(bucketName, objectKey, filePath string) error {
	result, err := c.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return fmt.Errorf("failed to get object, %v", err)
	}
	defer result.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s, %v", filePath, err)
	}
	defer file.Close()

	_, err = io.Copy(file, result.Body)
	if err != nil {
		return fmt.Errorf("failed to write file, %v", err)
	}

	return nil
}

func (c *s3Client) DeleteFile(bucketName, objectKey string) error {
	_, err := c.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return fmt.Errorf("failed to delete object, %v", err)
	}

	return nil
}
