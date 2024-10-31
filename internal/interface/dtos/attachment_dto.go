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

type GetAllAttachmentsDTO struct {
	Page         	int
	PageSize     	int
	DisplayName  	string
	AttachmentType 	string // type: docs, image
	StartTime    	time.Time
	EndTime      	time.Time
}

type GetAllAttachmentsByRoleDTO struct {
	Page         	int
	PageSize     	int
	DisplayName  	string
	AttachmentType 	string // type: docs, image
	Role			string
	StartTime    	time.Time
	EndTime      	time.Time
}
