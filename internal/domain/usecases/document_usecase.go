package usecases

import (
	"errors"
	"fmt"

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

func (u *documentUsecase) GetAllDocuments() (*[]dtos.DocumentDTO, *apperror.AppError) {
	return nil, nil
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
	existingDocument, err := u.documentRepository.FindDocumentByID(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Named("UpdateDocumentByID").Error(constant.ErrDocumentNotFound, zap.String("documentID", ID))
			return apperror.InternalServerError(constant.ErrDocumentNotFound)
		}
		u.logger.Named("UpdateDocumentByID").Error(constant.ErrFindDocumentByID, zap.String("documentID", ID), zap.Error(err))
		return apperror.InternalServerError(constant.ErrFindDocumentByID)
	}

	if err := u.documentRepository.UpdateDocumentByID(ID, updateMap); err != nil {
		u.logger.Named("UpdateDocumentByID").Error(constant.ErrUpdateDocumentFailed, zap.String("documentID", ID), zap.Error(err))
		return apperror.InternalServerError(constant.ErrUpdateDocumentFailed)
	}

	u.logger.Named("UpdateDocumentByID").Info("Success: Document updated", zap.String("documentID", existingDocument.ID))
	return nil
}

func (u *documentUsecase) DeleteDocumentByID(ID string) *apperror.AppError {
	existingDocument, err := u.documentRepository.FindDocumentByID(ID)
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Named("DeleteDocumentByID").Error(constant.ErrDocumentNotFound, zap.String("documentID", ID))
			return apperror.InternalServerError(constant.ErrDocumentNotFound)
		}
		u.logger.Named("DeleteDocumentByID").Error(constant.ErrFindDocumentByID, zap.String("documentID", ID), zap.Error(err))
		return apperror.InternalServerError(constant.ErrFindDocumentByID)
	}
	if err := u.documentRepository.DeleteDocumentByID(ID); err != nil {
		u.logger.Named("DeleteDocumentByID").Error(constant.ErrDeleteDocumentFailed, zap.String("documentID", ID), zap.Error(err))
		return apperror.InternalServerError(constant.ErrDeleteDocumentFailed)
	}

	u.logger.Named("DeleteDocumentByID").Info("Success: Document deleted", zap.String("documentID", existingDocument.ID))
	return nil
}
