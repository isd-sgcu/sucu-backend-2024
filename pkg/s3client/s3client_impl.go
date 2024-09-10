package s3client

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Client struct {
	client *s3.Client
}

func NewS3Client(cfg config.Config) S3Client {
	// Get credentials from env variable
	accessKey := cfg.GetAws().AccessKeyId
	secretKey := cfg.GetAws().SecretAccessKey
	region := cfg.GetAws().Region

	if accessKey == "" || secretKey == "" || region == "" {
		panic("AWS credentials or region not found in .env file")
	}

	// Create the credentials object
	creds := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")

	// Create the S3 client
	client := s3.NewFromConfig(aws.Config{
		Region:      region,
		Credentials: creds,
	})

	return &s3Client{
		client: client,
	}
}

func (c *s3Client) UploadFile(bucketName, objectKey string, file io.Reader) error {
	_, err := c.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}

	return nil
}

func (c *s3Client) DownloadFile(bucketName, objectKey, filePath string) error {
	result, err := c.client.GetObject(context.TODO(), &s3.GetObjectInput{
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
	_, err := c.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return fmt.Errorf("failed to delete object, %v", err)
	}

	return nil
}
