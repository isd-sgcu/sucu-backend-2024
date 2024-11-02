package repositories

import (
	"io"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
)

type AttachmentRepository interface {
	// client side
	FindAllAttachments() (*entities.Attachment, error)

	// back office
	FindAttachmentByID(ID string) (*entities.Attachment, error)

	InsertAttachments(attachments *[]entities.Attachment) error
	UploadAttachmentToS3(bucketName string, fileReaders map[string]io.Reader) error

	DeleteAttachmentByID(ID string) error
	DeleteAttachmentFromS3(bucketName, objectKey string) error
}
