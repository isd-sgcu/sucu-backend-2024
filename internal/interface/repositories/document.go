package repositories

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
)

type DocumentRepository interface {
	// client side
	FindAllDocuments(args *dtos.FindAllDocumentsDTO) (*[]entities.Document, error)
	FindDocumentByID(ID string) (*entities.Document, error)

	// back office
	FindDocumentsByRole(roles *[]string) (*[]entities.Document, error)
	InsertDocument(document *entities.Document) error
	UpdateDocumentByID(ID string, updateMap interface{}) error
	DeleteUserByID(ID string) error
}
