package usecases

import (
	"errors"

	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"
	"github.com/isd-sgcu/sucu-backend-2024/utils/constant"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type documentUsecase struct {
	cfg                config.Config
	logger             *zap.Logger
	documentRepository repositories.DocumentRepository
}

func NewDocumentUsecase(cfg config.Config, logger *zap.Logger, documentRepository repositories.DocumentRepository) DocumentUsecase {
	return &documentUsecase{
		cfg:                cfg,
		logger:             logger,
		documentRepository: documentRepository,
	}
}

func (u *documentUsecase) GetAllDocuments() (*[]dtos.DocumentDTO, error) {
	return nil, nil
}

func (u *documentUsecase) GetDocumentByID(ID string) (*dtos.DocumentDTO, error) {
	return nil, nil
}

func (u *documentUsecase) GetDocumentsByRole(req *dtos.UserDTO) (*[]dtos.DocumentDTO, error) {
	return nil, nil
}

func (u *documentUsecase) CreateDocument(document *dtos.CreateDocumentDTO) error {
	return nil
}

func (u *documentUsecase) UpdateDocumentByID(ID string, updateMap interface{}) error {
	existingDocument, err := u.documentRepository.FindDocumentByID(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Named("UpdateDocumentByID").Error(constant.ErrDocumentNotFound.Error(), zap.String("documentID", ID))
			return constant.ErrDocumentNotFound
		}
		u.logger.Named("UpdateDocumentByID").Error(constant.ErrFindDocumentByID.Error(), zap.String("documentID", ID), zap.Error(err))
		return constant.ErrFindDocumentByID
	}

	// Update the document with the provided updateMap
	if err := u.documentRepository.UpdateDocumentByID(ID, updateMap); err != nil {
		u.logger.Named("UpdateDocumentByID").Error(constant.ErrUpdateDocumentFailed.Error(), zap.String("documentID", ID), zap.Error(err))
		return constant.ErrUpdateDocumentFailed
	}

	u.logger.Named("UpdateDocumentByID").Info("Success: Document updated", zap.String("documentID", existingDocument.ID))
	return nil
}

func (u *documentUsecase) DeleteDocumentByID(ID string) error {
	existingDocument, err := u.documentRepository.FindDocumentByID(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Named("DeleteDocumentByID").Error(constant.ErrDocumentNotFound.Error(), zap.String("documentID", ID))
			return constant.ErrDocumentNotFound
		}
		u.logger.Named("DeleteDocumentByID").Error(constant.ErrFindDocumentByID.Error(), zap.String("documentID", ID), zap.Error(err))
		return constant.ErrFindDocumentByID
	}

	// Delete the document
	if err := u.documentRepository.DeleteDocumentByID(ID); err != nil {
		u.logger.Named("DeleteDocumentByID").Error(constant.ErrDeleteDocumentFailed.Error(), zap.String("documentID", ID), zap.Error(err))
		return constant.ErrDeleteDocumentFailed
	}

	u.logger.Named("DeleteDocumentByID").Info("Success: Document deleted", zap.String("documentID", existingDocument.ID))
	return nil
}
