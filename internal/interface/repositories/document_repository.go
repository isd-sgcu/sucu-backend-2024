package repositories

import (
	"time"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
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

type FindAllDocumentsArgs struct {
	Offset int
	Limit int
	DocumentType string
	Organization string
	Query string
	StartTime time.Time
	EndTime time.Time
}

// client side
func (r *documentRepository) FindAllDocuments(args FindAllDocumentsArgs) (*[]entities.Document, error) {
	var documents []entities.Document

	err := r.db.Joins("user").Where(`documents.type LIKE %?% 
									AND documents.content LIKE %?% 
									AND documents.createdAt BETWEEN ? AND ? 
									AND users.role_id LIKE %?%`, 
									args.DocumentType, 
									args.Query, 
									args.StartTime, 
									args.EndTime, 
									args.Organization).Offset(args.Offset).Limit(args.Limit).Find(&documents).Error	
	
	if err != nil {
		return nil, err
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
