package usecases

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/apperror"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"
	"github.com/isd-sgcu/sucu-backend-2024/utils"
	"github.com/isd-sgcu/sucu-backend-2024/utils/constant"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type documentUsecase struct {
	cfg                config.Config
	logger             *zap.Logger
	documentRepository repositories.DocumentRepository
	userRepository     repositories.UserRepository
}

func NewDocumentUsecase(cfg config.Config, logger *zap.Logger, documentRepository repositories.DocumentRepository, userRepository repositories.UserRepository) DocumentUsecase {
	return &documentUsecase{
		cfg:                cfg,
		logger:             logger,
		documentRepository: documentRepository,
		userRepository:     userRepository,
	}
}

func (u *documentUsecase) GetAllDocuments(req *dtos.GetAllDocumentsDTO) (*dtos.PaginationResponse, *apperror.AppError) {
	// retreive documents from repository
	args := &repositories.FindAllDocumentsArgs{
		Offset:       (req.Page - 1) * req.PageSize,
		Limit:        req.PageSize,
		DocumentType: req.DocumentType,
		Organization: req.Organization,
		Title:        req.Title,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
	}

	documents, err := u.documentRepository.FindAllDocuments(args)
	if err != nil {
		u.logger.Named("GetAllDocuments").Error(constant.ErrGetDocumentFailed, zap.Error(err))
		return nil, apperror.InternalServerError(constant.ErrGetDocumentFailed)
	}

	// create pagination response dtos
	data := make([]map[string]interface{}, 0)
	for _, d := range *documents {
		data = append(data, map[string]interface{}{
			"id":           d.ID,
			"title":        d.Title,
			"banner":       d.Banner,
			"cover":        d.Cover,
			"type":         strings.ToLower(d.TypeID),
			"created_at":   d.CreatedAt,
			"updated_at":   d.UpdatedAt,
			"organization": strings.ToLower(strings.Split(d.Author.RoleID, "_")[0]),
		})
	}

	paginationResponse := dtos.PaginationResponse{
		Data:      data,
		Page:      fmt.Sprintf("%d", req.Page),
		Limit:     fmt.Sprintf("%d", req.PageSize),
		TotalPage: fmt.Sprintf("%d", (int(math.Ceil(float64(len(data)) / float64(req.PageSize))))),
	}

	return &paginationResponse, nil
}

func (u *documentUsecase) GetDocumentByID(ID string) (*dtos.DocumentDTO, *apperror.AppError) {
	document, err := u.documentRepository.FindDocumentByID(ID)
	if err != nil {
		u.logger.Named("GetDocumentByID").Error(constant.ErrGetDocumentFailed, zap.String("documentID", ID), zap.Error(err))
		return nil, apperror.NotFoundError(constant.ErrDocumentNotFound)
	}
	resReturn := dtos.DocumentDTO{
		ID:        document.ID,
		Title:     document.Title,
		Banner:    document.Banner,
		Cover:     document.Cover,
		Content:   document.Content,
		UserID:    document.UserID,
		TypeID:    document.TypeID,
		CreatedAt: document.CreatedAt,
		UpdatedAt: document.UpdatedAt,
	}

	return &resReturn, nil
}

func (u *documentUsecase) GetDocumentsByRole(req *dtos.GetAllDocumentsByRoleDTO) (*dtos.PaginationResponse, *apperror.AppError) {
	// retreive documents from repository
	args := &repositories.FindAllDocumentsByRoleArgs{
		Offset:       (req.Page - 1) * req.PageSize,
		Limit:        req.PageSize,
		DocumentType: req.DocumentType,
		Organization: req.Organization,
		Title:        req.Title,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		Role:         req.Role,
	}

	documents, err := u.documentRepository.FindDocumentsByRole(args)
	if err != nil {
		u.logger.Named("GetAllDocumentsByRole").Error(constant.ErrGetDocumentFailed, zap.Error(err))
		return nil, apperror.InternalServerError(constant.ErrGetDocumentFailed)
	}

	// create pagination response dtos
	data := make([]map[string]interface{}, 0)
	for _, d := range *documents {
		data = append(data, map[string]interface{}{
			"id":           d.ID,
			"title":        d.Title,
			"banner":       d.Banner,
			"cover":        d.Cover,
			"type":         strings.ToLower(d.TypeID),
			"created_at":   d.CreatedAt,
			"updated_at":   d.UpdatedAt,
			"organization": strings.ToLower(strings.Split(d.Author.RoleID, "_")[0]),
			"author_role":  strings.ToLower(d.Author.RoleID),
		})
	}

	paginationResponse := dtos.PaginationResponse{
		Data:      data,
		Page:      fmt.Sprintf("%d", req.Page),
		Limit:     fmt.Sprintf("%d", req.PageSize),
		TotalPage: fmt.Sprintf("%d", (int(math.Ceil(float64(len(data)) / float64(req.PageSize))))),
	}

	return &paginationResponse, nil
}

func (u *documentUsecase) CreateDocument(document *dtos.CreateDocumentDTO) *apperror.AppError {
	// validate user
	_, err := u.userRepository.FindUserByID(document.UserID)
	if err != nil {
		u.logger.Named("CreateDocument").Error(constant.ErrUserNotFound, zap.String("user_id", document.UserID), zap.Error(err))
		return apperror.NotFoundError(constant.ErrUserNotFound)
	}

	docType, err := utils.GetDocType(document.TypeID)
	if err != nil {
		u.logger.Named("CreateDocument").Error(constant.ErrInvalidDocType, zap.String("type_id", document.TypeID), zap.Error(err))
		return apperror.BadRequestError(constant.ErrInvalidDocType)
	}

	newDocument := &entities.Document{
		ID:      fmt.Sprintf("DOC-%v", utils.GenerateRandomString("0123456789", 8)),
		Title:   document.Title,
		Content: document.Content,
		Banner:  document.Banner,
		Cover:   document.Cover,
		UserID:  document.UserID,
		TypeID:  docType,
	}

	if err := u.documentRepository.InsertDocument(newDocument); err != nil {
		u.logger.Named("CreateDocument").Error(constant.ErrInsertDocumentFailed, zap.String("document_id", newDocument.ID), zap.Error(err))
		return apperror.InternalServerError(constant.ErrInsertDocumentFailed)
	}

	u.logger.Named("CreateDocument").Info("Success: ", zap.String("document_id", newDocument.ID))
	return nil
}

func (u *documentUsecase) UpdateDocumentByID(ID string, updateMap interface{}) *apperror.AppError {
	if err := u.documentRepository.UpdateDocumentByID(ID, updateMap); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Named("UpdateDocumentByID").Error(constant.ErrDocumentNotFound, zap.String("documentID", ID))
			return apperror.InternalServerError(constant.ErrDocumentNotFound)
		}
		u.logger.Named("UpdateDocumentByID").Error(constant.ErrUpdateDocumentFailed, zap.String("documentID", ID), zap.Error(err))
		return apperror.InternalServerError(constant.ErrUpdateDocumentFailed)
	}

	u.logger.Named("UpdateDocumentByID").Info("Success: Document updated", zap.String("documentID", ID))
	return nil
}

func (u *documentUsecase) DeleteDocumentByID(ID string) *apperror.AppError {
	if err := u.documentRepository.DeleteDocumentByID(ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Named("DeleteDocumentByID").Error(constant.ErrDocumentNotFound, zap.String("documentID", ID))
			return apperror.InternalServerError(constant.ErrDocumentNotFound)
		}
		u.logger.Named("DeleteDocumentByID").Error(constant.ErrDeleteDocumentFailed, zap.String("documentID", ID), zap.Error(err))
		return apperror.InternalServerError(constant.ErrDeleteDocumentFailed)
	}

	u.logger.Named("DeleteDocumentByID").Info("Success: Document deleted", zap.String("documentID", ID))
	return nil
}
