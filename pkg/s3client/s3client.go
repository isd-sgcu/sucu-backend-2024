package s3client

import "io"

type S3Client interface {
	UploadFile(bucketName, objectKey string, file io.Reader) error
	DownloadFile(bucketName, objectKey, filePath string) error
	DeleteFile(bucketName, objectKey string) error
}
