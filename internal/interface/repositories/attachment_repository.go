package repositories

import (
	"io"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/s3client"
	"gorm.io/gorm"
)

type attachmentRepository struct {
	db *gorm.DB
	s3 s3client.S3Client
}

func NewAttachmentRepository(db *gorm.DB, s3 s3client.S3Client) AttachmentRepository {
	return &attachmentRepository{
		db: db,
		s3: s3,
	}
}

// client side
func (r *attachmentRepository) FindAllAttachments() (*entities.Attachment, error) {
	return nil, nil
}

// back office
func (r *attachmentRepository) InsertAttachments(attachments *[]entities.Attachment) error {
	if err := r.db.Create(&attachments).Error; err != nil {
		return err
	}

	return nil
}

func (r *attachmentRepository) UploadAttachmentToS3(bucketName string, fileReaders map[string]io.Reader) error {
	for fileName, file := range fileReaders {
		if err := r.s3.UploadFile(bucketName, fileName, file); err != nil {
			return err
		}
	}

	return nil
}

func (r *attachmentRepository) DeleteAttachmentByID(ID string) error {
	return nil
}

func (r *attachmentRepository) DeleteAttachmentFromS3(bucketName, objectKey string) error {
	return nil
}
