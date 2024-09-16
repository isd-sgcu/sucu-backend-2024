package s3client

import "bytes"

type S3Client interface {
	UploadFile(bucketName, objectKey string, buffer *bytes.Reader) error
	DownloadFile(bucketName, objectKey, filePath string) error
	DeleteFile(bucketName, objectKey string) error
}
