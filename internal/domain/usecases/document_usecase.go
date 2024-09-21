package usecases

import (
	"fmt"

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

func (u *documentUsecase) GetAllDocuments() (*[]dtos.DocumentDTO, *apperror.AppError) {
	return nil, nil
}

func (u *documentUsecase) GetDocumentByID(ID string) (*dtos.DocumentDTO, *apperror.AppError) {
	return nil, nil
}

func (u *documentUsecase) GetDocumentsByRole(req *dtos.UserDTO) (*[]dtos.DocumentDTO, *apperror.AppError) {
	return nil, nil
}

func (u *documentUsecase) CreateDocument(document *dtos.CreateDocumentDTO) *apperror.AppError { // TODO: implement this
	// validate document
	if document.Title == "" || document.Content == "" || document.UserID == "" || document.TypeID == "" {
		u.logger.Named("CreateDocument").Error("Invalid document", zap.String("title", document.Title), zap.String("content", document.Content), zap.String("user_id", document.UserID), zap.String("type_id", document.TypeID))
		return apperror.BadRequestError("invalid document")
	}

	// validate user
	existingUser, err := u.userRepository.FindUserByID(document.UserID)
	if err != nil {
		u.logger.Named("CreateDocument").Error(constant.ErrFindUserByID, zap.String("user_id", document.UserID), zap.Error(err))
		return apperror.InternalServerError(constant.ErrFindUserByID)
	}
	if existingUser == nil {
		u.logger.Named("CreateDocument").Error(constant.ErrUserNotFound, zap.String("user_id", document.UserID))
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
