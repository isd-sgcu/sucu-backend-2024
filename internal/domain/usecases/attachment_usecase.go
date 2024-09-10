package usecases

import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"
	"github.com/isd-sgcu/sucu-backend-2024/utils"
	"github.com/isd-sgcu/sucu-backend-2024/utils/constant"
	"go.uber.org/zap"
)

type attachmentUsecase struct {
	cfg                  config.Config
	logger               *zap.Logger
	attachmentRepository repositories.AttachmentRepository
}

func NewAttachmentUsecase(cfg config.Config, logger *zap.Logger, attachmentRepository repositories.AttachmentRepository) AttachmentUsecase {
	return &attachmentUsecase{
		cfg:                  cfg,
		logger:               logger,
		attachmentRepository: attachmentRepository,
	}
}

func (u *attachmentUsecase) GetAllAttachments() (*[]dtos.AttachmentDTO, error) {
	return nil, nil
}

func (u *attachmentUsecase) GetAllAttachmentsByRole(req dtos.UserDTO) (*[]dtos.AttachmentDTO, error) {
	return nil, nil
}

func (u *attachmentUsecase) CreateAttachments(documentID string, files map[string][]*multipart.FileHeader) error {
	fileReaders := make(map[string]io.Reader)

	for _, fileHeaders := range files {
		for _, fileHeader := range fileHeaders {
			// Check the file size
			if fileHeader.Size > constant.MAX_FILE_SIZE {
				u.logger.Named("CreateAttachments").Error("Check the file size: ", zap.Error(fmt.Errorf("file size exceeds the allowed limit")))
				return fmt.Errorf("file size exceeds the allowed limit")
			}

			// Open the uploaded file
			src, err := fileHeader.Open()
			if err != nil {
				u.logger.Named("CreateAttachments").Error("Open uploaded file: ", zap.Error(err))
				return fmt.Errorf("failed to open the uploaded file: %w", err)
			}
			defer src.Close()

			// Validate the file type
			_, err = utils.ValidateFileType(fileHeader.Filename)
			if err != nil {
				u.logger.Named("CreateAttachments").Error("Validate file type: ", zap.Error(err))
				return fmt.Errorf("invalid file type for %s: %w", fileHeader.Filename, err)
			}

			// Add the file reader to the map
			fileReaders[uuid.New().String()] = src
		}
	}

	if err := u.attachmentRepository.UploadAttachmentToS3(u.cfg.GetAws().BucketName, fileReaders); err != nil {
		u.logger.Error(err.Error())
		return err
	}

	var attachments []entities.Attachment
	for fileName := range fileReaders {
		attachment := entities.Attachment{
			ID:         fileName,
			Path:       fileName,
			DocumentID: documentID,
			TypeID: func() string {
				typeID, _ := utils.ValidateFileType(fileName)
				return *typeID
			}(),
		}
		attachments = append(attachments, attachment)
	}
	if err := u.attachmentRepository.InsertAttachments(&attachments); err != nil {
		return err
	}

	return nil
}

func (u *attachmentUsecase) DeleteAttachment(ID string) error {
	return nil
}
