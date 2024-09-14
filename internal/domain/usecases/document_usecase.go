package usecases

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"
	"github.com/isd-sgcu/sucu-backend-2024/utils"
	"github.com/isd-sgcu/sucu-backend-2024/utils/dtoconv"
	"go.uber.org/zap"
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

func (u *documentUsecase) GetAllDocuments(req *dtos.GetDocumentsDTO) (*[]dtos.DocumentDTO, error) {
	findAllDocumentsDTO := dtos.FindAllDocumentsDTO{
		Page: req.Page,
		Limit: req.Limit,
		Query: req.Query,
		Org: req.Org,
		Type: req.Type,
	}

	documents, err := u.documentRepository.FindAllDocuments(&findAllDocumentsDTO)
	if err != nil {
		u.logger.Named("GetAllDocuments").Error("get all documents: ", zap.Error(err))
		return nil, err
	}

	var documentsDTO = make([]dtos.DocumentDTO, len(*documents))
	for i , document := range *documents {
		documentsDTO[i] = *dtoconv.ConvertToDocumentDTO(&document)
	}
	
	return &documentsDTO, nil
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
	return nil
}

func (u *documentUsecase) DeleteDocumentByID(ID string) error {
	return nil
}
