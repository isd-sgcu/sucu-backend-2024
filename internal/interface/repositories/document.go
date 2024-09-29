package repositories

import "github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"

type DocumentRepository interface {
	// client side
	FindAllDocuments() (*[]entities.Document, error)
	FindDocumentByID(ID string) (*entities.Document, error)

	// back office
	FindDocumentsByRole(roles *[]string) (*[]entities.Document, error)
	InsertDocument(document *entities.Document) error
	UpdateDocumentByID(ID string, updateMap interface{}) error
	DeleteDocumentByID(ID string) error
}
