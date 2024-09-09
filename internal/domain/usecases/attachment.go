package usecases

import "mime/multipart"

type AttachmentUsecase interface {
	CreateAttachments(documentID string, files map[string][]*multipart.FileHeader) error
}
