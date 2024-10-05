package repositories

import (
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
	Title        string
	StartTime    time.Time
	EndTime      time.Time
}

// client side
func (r *documentRepository) FindAllDocuments(args *FindAllDocumentsArgs) (*[]entities.Document, error) {
	var documents []entities.Document

	r.db.Raw(`
		SELECT * FROM documents INNER JOIN users ON documents.user_id = users.id
		WHERE documents.type_id LIKE %?% 
		AND	 LOWER(documents.title) LIKE %?%
		AND  documents.created_at BETWEEN %?% AND %?%
		OFFSET ? LIMIT ?`,
		strings.ToUpper(args.DocumentType),
		strings.ToLower(args.Title),
		args.StartTime,
		args.EndTime,
		args.Offset,
		args.Limit).Scan(&documents)

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
	return r.db.Create(document).Error
}

func (r *documentRepository) UpdateDocumentByID(ID string, updateMap interface{}) error {
	return nil
}

func (r *documentRepository) DeleteUserByID(ID string) error {
	return nil
}
