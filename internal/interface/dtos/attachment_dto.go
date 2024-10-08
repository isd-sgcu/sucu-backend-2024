package dtos

import "time"

type AttachmentDTO struct {
	ID          string    `json:"id"`
	DisplayName string    `json:"name"`
	DocumentID  string    `json:"document_id"`
	TypeID      string    `json:"type_id"`
	RoleID      string    `json:"role_id"` // role_id จะเอาไว้ให้ client ดูว่าไฟล์นี้มาจาก org อะไร sgcu or sucu
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
