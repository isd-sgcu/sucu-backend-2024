package dtos

import "time"

type DocumentDTO struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Banner    *string   `json:"banner"`
	Cover     *string   `json:"cover"`
	UserID    string    `json:"user_id"`
	TypeID    string    `json:"type_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Author UserDTO         `json:"author"`
	Images []AttachmentDTO `json:"images"` // images file eg. jpeg jpg png
	Docs   []AttachmentDTO `json:"docs"`   // docs file eg. pdf xlsx pptx
}

type CreateDocumentDTO struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Banner  *string `json:"banner"`
	Cover   *string `json:"cover"`
	UserID  string  `json:"user_id"`
	TypeID  string  `json:"type_id"`
}

type UpdateDocumentDTO struct {
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Banner  *string `json:"banner"`
	Cover   *string `json:"cover"`
}
