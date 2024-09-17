package dtoconv

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
)

func ConvertToDocumentDTO(document *entities.Document) *dtos.DocumentDTO {
	// Create a new DocumentDTO instance
	documentDTO := &dtos.DocumentDTO{
		ID:        document.ID,
		Title:     document.Title,
		Content:   document.Content,
		Banner:    document.Banner,
		Cover:     document.Cover,
		UserID:    document.UserID,
		TypeID:    document.TypeID,
		CreatedAt: document.CreatedAt,
		UpdatedAt: document.UpdatedAt,
	}

	// Convert the Author field from User to UserDTO
	documentDTO.Author = dtos.UserDTO{
		ID:        document.Author.ID,
		FirstName: document.Author.FirstName,
		LastName:  document.Author.LastName,
		Role:      document.Author.RoleID,
		CreatedAt: document.Author.CreatedAt,
		UpdatedAt: document.Author.UpdatedAt,
	}

	// Convert the Images and docs field from []Attachment to []AttachmentDTO
	// check organization by author's role
	var org string
	role := document.Author.RoleID
	if role == "SGCU_ADMIN" || role == "SGCU_SUPERADMIN" {
		org = "sgcu"
	} else {
		org = "sccu"
	}

	for _, attachment := range document.Attachments {
		if attachment.TypeID == "image" {
			documentDTO.Images = append(documentDTO.Images, dtos.AttachmentDTO{
				ID:          attachment.ID,
				DisplayName: attachment.DisplayName,
				DocumentID:  attachment.DocumentID,
				TypeID:      attachment.TypeID,
				RoleID:      org,
				CreatedAt:   attachment.CreatedAt,
				UpdatedAt:   attachment.UpdatedAt,
			})
		} else if attachment.TypeID == "doc" {
			documentDTO.Docs = append(documentDTO.Docs, dtos.AttachmentDTO{
				ID:          attachment.ID,
				DisplayName: attachment.DisplayName,
				DocumentID:  attachment.DocumentID,
				TypeID:      attachment.TypeID,
				RoleID:      org,
				CreatedAt:   attachment.CreatedAt,
				UpdatedAt:   attachment.UpdatedAt,
			})
		}
	}

	return documentDTO
}