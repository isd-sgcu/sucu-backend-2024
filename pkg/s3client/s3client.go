package s3client

type S3Client interface {
	UploadFile(bucketName, objectKey, filePath string) error
	DownloadFile(bucketName, objectKey, filePath string) error
	DeleteObject(bucketName, objectKey string) error
}
