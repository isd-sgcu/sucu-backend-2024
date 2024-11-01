package repositories

import (
	"io"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/s3client"
	"github.com/isd-sgcu/sucu-backend-2024/utils"
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
func (r *attachmentRepository) FindAttachmentByID(ID string) (*entities.Attachment, error) {
	var attachment entities.Attachment
	if err := r.db.First(&attachment, "id = ?", ID).Error; err != nil {
		return nil, err
	}
	return &attachment, nil
}

func (r *attachmentRepository) InsertAttachments(attachments *[]entities.Attachment) error {
	if err := r.db.Create(&attachments).Error; err != nil {
		return err
	}

	return nil
}

func (r *attachmentRepository) UploadAttachmentToS3(bucketName string, fileReaders map[string]io.Reader) error {
	for fileName, file := range fileReaders {
		buffer, err := utils.ToBytesReader(file)
		if err != nil {
			return err
		}
		if err := r.s3.UploadFile(bucketName, fileName, buffer); err != nil {
			return err
		}
	}

	return nil
}

func (r *attachmentRepository) DeleteAttachmentByID(ID string) error {
	if err := r.db.Delete(&entities.Attachment{}, "id = ?", ID).Error; err != nil {
		return err
	}
	return nil
}

func (r *attachmentRepository) DeleteAttachmentFromS3(bucketName, objectKey string) error {
	return nil
}
