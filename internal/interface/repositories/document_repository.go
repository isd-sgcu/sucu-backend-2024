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
		WHERE users.role_id LIKE ?
		AND  documents.type_id LIKE ?
		AND	 LOWER(documents.title) LIKE ?
		AND  documents.created_at BETWEEN ? AND ?
		OFFSET ? LIMIT ?`,
		fmt.Sprintf("%%%s%%", strings.ToUpper(args.Organization)),
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

type FindAllDocumentsByRoleArgs struct {
	Offset       int
	Limit        int
	DocumentType string
	Organization string
	Title        string
	StartTime    time.Time
	EndTime      time.Time
	Role         string
}

// back office
func (r *documentRepository) FindDocumentsByRole(args *FindAllDocumentsByRoleArgs) (*[]entities.Document, error) {
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
		AND  users.role_id LIKE ?
		AND  users.role_id = ?
		AND  documents.created_at BETWEEN ? AND ?
		OFFSET ? LIMIT ?`,
		fmt.Sprintf("%%%s%%", strings.ToUpper(args.DocumentType)),
		fmt.Sprintf("%%%s%%", strings.ToLower(args.Title)),
		fmt.Sprintf("%%%s%%", strings.ToUpper(args.Organization)),
		strings.ToUpper(args.Role),
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

func (r *documentRepository) InsertDocument(document *entities.Document) error {
	return r.db.Create(document).Error
}

func (r *documentRepository) UpdateDocumentByID(ID string, updateMap interface{}) error {
	result := r.db.Model(&entities.Document{}).Where("id = ?", ID).Updates(updateMap)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *documentRepository) DeleteDocumentByID(ID string) error {
	result := r.db.Where("id = ?", ID).Delete(&entities.Document{})
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
