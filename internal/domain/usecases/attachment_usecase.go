package usecases

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

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
	var attachments []entities.Attachment
	fileReaders := make(map[string]io.Reader)

	for _, fileHeaders := range files {
		for _, fileHeader := range fileHeaders {
			if err := u.validateAndProcessFile(fileHeader, documentID, &attachments, fileReaders); err != nil {
				u.logger.Named("CreateAttachments").Error("Validate and process file: ", zap.Error(err))
				return err
			}
		}
	}

	if err := u.uploadAndSaveAttachments(fileReaders, attachments); err != nil {
		u.logger.Named("CreateAttachments").Error("Upload and save attachments: ", zap.Error(err))
		return err
	}

	u.logger.Named("CreateAttachments").Info("Success: ", zap.String("document_id", documentID), zap.Any("files", fileReaders))
	return nil
}

func (u *attachmentUsecase) validateAndProcessFile(fileHeader *multipart.FileHeader, documentID string, attachments *[]entities.Attachment, fileReaders map[string]io.Reader) error {
	if fileHeader.Size > constant.MAX_FILE_SIZE {
		return fmt.Errorf("file size exceeds the allowed limit")
	}

	src, err := fileHeader.Open()
	if err != nil {
		u.logger.Named("CreateAttachments").Error("Open uploaded file: ", zap.Error(err))
		return fmt.Errorf("failed to open the uploaded file: %w", err)
	}
	defer src.Close()

	typeID, err := utils.ValidateFileType(fileHeader.Filename)
	if err != nil {
		u.logger.Named("CreateAttachments").Error("Validate file type: ", zap.Error(err))
		return fmt.Errorf("invalid file type for %s: %w", fileHeader.Filename, err)
	}

	fileName := fileHeader.Filename
	name, err := u.generateNewFileName(fileName)
	if err != nil {
		return err
	}

	attachment := entities.Attachment{
		ID:          name,
		DisplayName: fileName,
		DocumentID:  documentID,
		TypeID:      *typeID,
	}
	*attachments = append(*attachments, attachment)

	fileReaders[name] = src

	return nil
}

func (u *attachmentUsecase) generateNewFileName(fileName string) (string, error) {
	lastDotIndex := strings.LastIndex(fileName, ".")
	if lastDotIndex == -1 {
		return "", errors.New("not found file extension")
	}
	nameWithoutExt := fileName[:lastDotIndex]
	ext := fileName[lastDotIndex+1:]
	randomString := utils.GenerateRandomString("0123456789", 8)
	return fmt.Sprintf("%s-%s.%s", nameWithoutExt, randomString, ext), nil
}

func (u *attachmentUsecase) uploadAndSaveAttachments(fileReaders map[string]io.Reader, attachments []entities.Attachment) error {
	if err := u.attachmentRepository.UploadAttachmentToS3(u.cfg.GetAws().BucketName, fileReaders); err != nil {
		u.logger.Named("CreateAttachments").Error("Upload attachment to s3", zap.Error(err))
		return err
	}

	if err := u.attachmentRepository.InsertAttachments(&attachments); err != nil {
		return err
	}

	return nil
}

func (u *attachmentUsecase) DeleteAttachment(ID string) error {
	return nil
}
