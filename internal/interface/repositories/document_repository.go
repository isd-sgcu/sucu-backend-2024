package repositories

import (
	"fmt"
	"strings"
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
	Offset       int
	Limit        int
	DocumentType string
	Organization string
	Query        string
	StartTime    time.Time
	EndTime      time.Time
}

// client side
func (r *documentRepository) FindAllDocuments(args *FindAllDocumentsArgs) (*[]entities.Document, error) {
	var documents []entities.Document

	query := r.db.Model(&entities.Document{}).
		Joins("INNER JOIN users ON documents.user_id = users.id").
		Where("users.role_id LIKE ?", fmt.Sprintf("%%%s%%", strings.ToUpper(args.Organization))).
		Where("documents.type_id LIKE ?", fmt.Sprintf("(?i)%%%s%%", strings.ToUpper(args.DocumentType))).
		Where("documents.title LIKE ?", fmt.Sprintf("%%%s%%", args.Query)).
		Where("documents.created_at BETWEEN ? AND ?", args.StartTime, args.EndTime).
		Offset(args.Offset).
		Limit(args.Limit)

	err := query.Find(&documents).Error
	if err != nil {
		return nil, err
	}

	return &documents, nil
	// return nil, nil

}

func (r *documentRepository) FindDocumentByID(ID string) (*entities.Document, error) {
	return nil, nil
}

// back office
func (r *documentRepository) FindDocumentsByRole(roles *[]string) (*[]entities.Document, error) {
	return nil, nil
}

func (r *documentRepository) InsertDocument(document *entities.Document) error {
	return r.db.Create(document).Error
}

func (r *documentRepository) UpdateDocumentByID(ID string, updateMap interface{}) error {
	return nil
}

func (r *documentRepository) DeleteUserByID(ID string) error {
	return nil
}
