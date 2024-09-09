package repositories

import (
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/s3client"

	"gorm.io/gorm"
)

type repository struct {
	UserRepository       UserRepository
	AttachmentRepository AttachmentRepository
}

func NewRepository(cfg config.Config, db *gorm.DB, s3 s3client.S3Client) Repository {
	return &repository{
		UserRepository:       NewUserRepository(db),
		AttachmentRepository: NewAttachmentRepository(db, s3),
	}
}

func (r *repository) User() UserRepository {
	return r.UserRepository
}

func (r *repository) Attachment() AttachmentRepository {
	return r.AttachmentRepository
}
