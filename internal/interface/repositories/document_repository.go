package repositories

import (
	"fmt"
	"strings"
	"time"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/jinzhu/copier"
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

	var results []struct {
		*entities.Document
		*entities.User
		DocumentID string
		AuthorID   string
	}

	err := r.db.Raw(`
		SELECT *, documents.id AS document_id, users.id AS AuthorID 
		FROM documents INNER JOIN users ON documents.user_id = users.id
		WHERE documents.type_id LIKE ?
		AND	 LOWER(documents.title) LIKE ?
		AND  documents.created_at BETWEEN ? AND ?
		OFFSET ? LIMIT ?`,
		fmt.Sprintf("%%%s%%", strings.ToUpper(args.DocumentType)),
		fmt.Sprintf("%%%s%%", strings.ToLower(args.Title)),
		args.StartTime,
		args.EndTime,
		args.Offset,
		args.Limit).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	for _, result := range results {
		var d entities.Document
		copier.Copy(&d, &result)
		copier.Copy(&d.Author, &result)

		d.ID = result.DocumentID
		d.Author.ID = result.AuthorID

		documents = append(documents, d)
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
	return r.db.Create(document).Error
}

func (r *documentRepository) UpdateDocumentByID(ID string, updateMap interface{}) error {
	return nil
}

func (r *documentRepository) DeleteUserByID(ID string) error {
	return nil
}
