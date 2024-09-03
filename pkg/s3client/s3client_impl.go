package s3client

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Client struct {
	client *s3.Client
}

func NewS3Client(cfg config.Config) (S3Client, error) {
	// Get credentials from env variable
	accessKey := cfg.GetAws().AwsAccessKeyId
	secretKey := cfg.GetAws().AwsSecretAccessKey
	region := cfg.GetAws().AwsRegion

	if accessKey == "" || secretKey == "" || region == "" {
		return nil, fmt.Errorf("AWS credentials or region not found in .env file")
	}

	// Create the credentials object
	creds := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")

	// Create the S3 client
	client := s3.NewFromConfig(aws.Config{
		Region:      region,
		Credentials: creds,
	})

	return &s3Client{client: client}, nil
}

func (c *s3Client) UploadFile(bucketName, objectKey, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s, %v", filePath, err)
	}
	defer file.Close()

	_, err = c.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}

	log.Printf("Successfully uploaded %s to %s/%s\n", filePath, bucketName, objectKey)
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

	log.Printf("Successfully downloaded %s/%s to %s\n", bucketName, objectKey, filePath)
	return nil
}

func (c *s3Client) DeleteObject(bucketName, objectKey string) error {
	_, err := c.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return fmt.Errorf("failed to delete object, %v", err)
	}

	log.Printf("Successfully deleted %s from %s\n", objectKey, bucketName)
	return nil
}
