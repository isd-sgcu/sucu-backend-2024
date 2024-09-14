package usecases

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
)

type DocumentUsecase interface {
	// client side
	GetAllDocuments(req *dtos.GetDocumentsDTO) (*[]dtos.DocumentDTO, error)
	GetDocumentByID(ID string) (*dtos.DocumentDTO, error)

	// back office
	GetDocumentsByRole(req *dtos.UserDTO) (*[]dtos.DocumentDTO, error)
	CreateDocument(document *dtos.CreateDocumentDTO) error
	UpdateDocumentByID(ID string, updateMap interface{}) error
	DeleteDocumentByID(ID string) error
}
