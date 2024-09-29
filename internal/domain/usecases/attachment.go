package usecases

import (
	"mime/multipart"

	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/apperror"
)

type AttachmentUsecase interface {
	// client side
	GetAllAttachments() (*[]dtos.AttachmentDTO, *apperror.AppError)

	// back office
	GetAllAttachmentsByRole(req dtos.UserDTO) (*[]dtos.AttachmentDTO, *apperror.AppError)
	CreateAttachments(documentID string, files map[string][]*multipart.FileHeader) *apperror.AppError
	DeleteAttachment(ID string) *apperror.AppError
}
