package usecases

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/apperror"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"
	"github.com/isd-sgcu/sucu-backend-2024/utils"
	"github.com/isd-sgcu/sucu-backend-2024/utils/constant"
	"go.uber.org/zap"
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
	// validate arguments
	if org := strings.ToLower(req.Organization); org != "" &&
		org != strings.ToLower(constant.SCCU) &&
		org != strings.ToLower(constant.SGCU) {

		u.logger.Named("GetAllDocuments").Error(constant.ErrInvalidOrg, zap.String("organization", org))
		return nil, apperror.BadRequestError(constant.ErrInvalidOrg)
	}

	if dt := strings.ToLower(req.DocumentType); dt != "" &&
		dt != strings.ToLower(constant.ANNOUNCEMENT) &&
		dt != strings.ToLower(constant.STATISTIC) &&
		dt != strings.ToLower(constant.BUDGET) {

		u.logger.Named("GetAllDocuments").Error(constant.ErrInvalidDocType, zap.String("document_type", dt))
		return nil, apperror.BadRequestError(constant.ErrInvalidDocType)
	}

	if ps := req.PageSize; ps > constant.MAX_PAGE_SIZE || ps < 0 {
		u.logger.Named("GetAllDocuments").Error(constant.ErrInvalidPageSize, zap.Int("page_size", ps))
		return nil, apperror.BadRequestError(constant.ErrInvalidPageSize)
	}

	startTime, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		u.logger.Named("GetAllDocuments").Error(constant.ErrInvalidTimeFormat, zap.String("start time", req.StartTime), zap.Error(err))
		return nil, apperror.BadRequestError(constant.ErrInvalidTimeFormat)
	}
	endTime, err := time.Parse(time.RFC3339, req.EndTime)
	if err != nil {
		u.logger.Named("GetAllDocuments").Error(constant.ErrInvalidTimeFormat, zap.String("end time", req.EndTime), zap.Error(err))
		return nil, apperror.BadRequestError(constant.ErrInvalidTimeFormat)
	}

	// retreive documents from repository
	args := &repositories.FindAllDocumentsArgs{
		Offset:       (req.Page - 1) * req.PageSize,
		Limit:        req.PageSize,
		DocumentType: req.DocumentType,
		Organization: req.Organization,
		Title:        req.Title,
		StartTime:    startTime,
		EndTime:      endTime,
	}

	documents, err := u.documentRepository.FindAllDocuments(args)
	if err != nil {
		u.logger.Named("GetAllDocuments").Error(constant.ErrFindAllDocuments, zap.Error(err))
		return nil, apperror.InternalServerError(constant.ErrFindAllDocuments)
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
	return nil, nil
}

func (u *documentUsecase) GetDocumentsByRole(req *dtos.UserDTO) (*[]dtos.DocumentDTO, *apperror.AppError) {
	return nil, nil
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
	return nil
}

func (u *documentUsecase) DeleteDocumentByID(ID string) *apperror.AppError {
	return nil
}
