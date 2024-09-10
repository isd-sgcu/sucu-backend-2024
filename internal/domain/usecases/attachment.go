package usecases

import (
	"mime/multipart"

	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
)

type AttachmentUsecase interface {
	// client side
	GetAllAttachments() (*[]dtos.AttachmentDTO, error)

	// back office
	GetAllAttachmentsByRole(req dtos.UserDTO) (*[]dtos.AttachmentDTO, error)
	CreateAttachments(documentID string, files map[string][]*multipart.FileHeader) error
	DeleteAttachment(ID string) error
}
