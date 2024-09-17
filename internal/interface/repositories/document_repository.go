package repositories

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"gorm.io/gorm"
)

type documentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) DocumentRepository {
	return &documentRepository{
		db: db,
	}
}


// client side
func (r *documentRepository) FindAllDocuments(args *dtos.FindAllDocumentsDTO) (*[]entities.Document, error) {
	var documents []entities.Document
	result := r.db.Offset(args.Page * args.Limit).Limit(args.Limit).Where("").Find(&documents)
	if result.Error != nil {
		return nil, result.Error
	}

	

	return &documents, nil
}

func (r *documentRepository) FindDocumentByID(ID string) (*entities.Document, error) {
	return nil, nil
}

// back office
func (r *documentRepository) FindDocumentsByRole(roles *[]string) (*[]entities.Document, error) {
	return nil, nil
}

func (r *documentRepository) InsertDocument(document *entities.Document) error {
	return nil
}

func (r *documentRepository) UpdateDocumentByID(ID string, updateMap interface{}) error {
	return nil
}

func (r *documentRepository) DeleteUserByID(ID string) error {
	return nil
}
