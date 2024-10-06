package usecases

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/apperror"
)

type DocumentUsecase interface {
	// client side
	GetAllDocuments(req *dtos.GetAllDocumentsDTO) (*dtos.PaginationResponse, *apperror.AppError)
	GetDocumentByID(ID string) (*dtos.DocumentDTO, *apperror.AppError)

	// back office
	GetDocumentsByRole(req *dtos.GetAllDocumentsByRoleDTO) (*dtos.PaginationResponse, *apperror.AppError)
	CreateDocument(document *dtos.CreateDocumentDTO) *apperror.AppError
	UpdateDocumentByID(ID string, updateMap interface{}) *apperror.AppError
	DeleteDocumentByID(ID string) *apperror.AppError
}
