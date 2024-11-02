package repositories

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/s3client"
	"github.com/isd-sgcu/sucu-backend-2024/utils"
	"github.com/jinzhu/copier"
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

type FindAllAttachmentsArgs struct {
	Offset       		int
	Limit        		int
	AttachmentType 	   	string
	DisplayName        	string
	StartTime    		time.Time
	EndTime      		time.Time
}

// client side
func (r *attachmentRepository) FindAllAttachments(args *FindAllAttachmentsArgs) (*[]entities.Attachment, error) {
	var attachments []entities.Attachment

	var results []struct {
		*entities.Attachment
		*entities.Document
		AttachmentID  string
		DocumentID    string
	}

	err := r.db.Raw(`
		SELECT *, attachments.id AS attachment_id
		FROM attachments INNER JOIN documents ON documents.id = attachments.document_id
		AND	 LOWER(attachments.display_name) LIKE ?
		AND  attachments.created_at BETWEEN ? AND ?
		AND attachments.type_id = 'DOCS'
		OFFSET ? LIMIT ?`,
		fmt.Sprintf("%%%s%%", strings.ToLower(args.DisplayName)),
		args.StartTime,
		args.EndTime,
		args.Offset,
		args.Limit).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	for _, result := range results {
		var d entities.Attachment
		copier.Copy(&d, &result)
		copier.Copy(&d.Document, &result)

		d.ID = result.AttachmentID
		d.Document.ID = result.DocumentID

		attachments = append(attachments, d)
	}

	return &attachments, nil
}

// back office
type FindAllAttachmentsByRoleArgs struct {
	Offset       		int
	Limit        		int
	AttachmentType 	   	string
	DisplayName        	string
	StartTime    		time.Time
	EndTime      		time.Time
	Role				string
}

func (r *attachmentRepository) FindAllAttachmentsByRole(args *FindAllAttachmentsByRoleArgs) (*[]entities.Attachment, error) {
	var attachments []entities.Attachment

	var results []struct {
		*entities.Attachment
		*entities.Document
		AttachmentID  string
		DocumentID    string
	}

	err := r.db.Raw(`
		SELECT *, attachments.id AS attachment_id
		FROM attachments INNER JOIN documents ON documents.id = attachments.document_id
		INNER JOIN users ON users.id = documents.user_id
		AND	 LOWER(attachments.display_name) LIKE ?
		AND  attachments.created_at BETWEEN ? AND ?
		AND users.role_id = ?
		AND attachments.type_id = 'DOCS'
		OFFSET ? LIMIT ?`,
		fmt.Sprintf("%%%s%%", strings.ToLower(args.DisplayName)),
		args.StartTime,
		args.EndTime,
		strings.ToUpper(args.Role),
		args.Offset,
		args.Limit).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	for _, result := range results {
		var d entities.Attachment
		copier.Copy(&d, &result)
		copier.Copy(&d.Document, &result)

		d.ID = result.AttachmentID
		d.Document.ID = result.DocumentID

		attachments = append(attachments, d)
	}

	return &attachments, nil
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
	return nil
}

func (r *attachmentRepository) DeleteAttachmentFromS3(bucketName, objectKey string) error {
	return nil
}
