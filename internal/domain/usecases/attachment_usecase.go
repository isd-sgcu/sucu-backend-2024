package usecases

import (
	"errors"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"strings"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/apperror"
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

func (u *attachmentUsecase) GetAllAttachments(req *dtos.GetAllAttachmentsDTO) (*dtos.PaginationResponse, *apperror.AppError) {

	args := &repositories.FindAllAttachmentsArgs{
		Offset:       	(req.Page - 1) * req.PageSize,
		Limit:        	req.PageSize,
		AttachmentType: req.AttachmentType,
		DisplayName:    req.DisplayName,
		StartTime:    	req.StartTime,
		EndTime:      	req.EndTime,
	}

	attachments, err := u.attachmentRepository.FindAllAttachments(args)
	if err != nil {
		u.logger.Named("GetAllAttachments").Error(constant.ErrGetAttachmentFailed, zap.Error(err))
		return nil, apperror.InternalServerError(constant.ErrGetAttachmentFailed)
	}

	// create pagination response dtos
	data := make([]map[string]interface{}, 0)
	for _, d := range *attachments {
		data = append(data, map[string]interface{}{
			"id":           d.ID,
			"display_name": d.DisplayName,
			"document_id":  d.Document.ID,
			"type":         strings.ToLower(d.TypeID),
			"created_at":   d.CreatedAt,
			"updated_at":   d.UpdatedAt,
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

func (u *attachmentUsecase) GetAllAttachmentsByRole(req dtos.UserDTO) (*[]dtos.AttachmentDTO, *apperror.AppError) {
	return nil, nil
}

func (u *attachmentUsecase) CreateAttachments(documentID string, files map[string][]*multipart.FileHeader) *apperror.AppError {
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

func (u *attachmentUsecase) validateAndProcessFile(fileHeader *multipart.FileHeader, documentID string, attachments *[]entities.Attachment, fileReaders map[string]io.Reader) *apperror.AppError {
	if fileHeader.Size > constant.MAX_FILE_SIZE {
		return apperror.BadRequestError("file size exceeds the allowed limit")
	}

	src, err := fileHeader.Open()
	if err != nil {
		u.logger.Named("CreateAttachments").Error("Open uploaded file: ", zap.Error(err))
		return apperror.InternalServerError(fmt.Sprintf("failed to open the uploaded file: %s", err.Error()))
	}
	defer src.Close()

	typeID, err := utils.ValidateFileType(fileHeader.Filename)
	if err != nil {
		u.logger.Named("CreateAttachments").Error("Validate file type: ", zap.Error(err))
		return apperror.BadRequestError(fmt.Sprintf("invalid file type for %s: %s", fileHeader.Filename, err.Error()))
	}

	fileName := fileHeader.Filename
	name, err := u.generateNewFileName(fileName)
	if err != nil {
		return apperror.InternalServerError(fmt.Sprintf("failed to generate new file name: %s", err.Error()))
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

func (u *attachmentUsecase) uploadAndSaveAttachments(fileReaders map[string]io.Reader, attachments []entities.Attachment) *apperror.AppError {
	if err := u.attachmentRepository.UploadAttachmentToS3(u.cfg.GetAws().BucketName, fileReaders); err != nil {
		u.logger.Named("CreateAttachments").Error("Upload attachment to s3", zap.Error(err))
		return apperror.InternalServerError(fmt.Sprintf("failed to upload attachment to s3: %s", err.Error()))
	}

	if err := u.attachmentRepository.InsertAttachments(&attachments); err != nil {
		return apperror.InternalServerError(fmt.Sprintf("failed to insert attachments: %s", err.Error()))
	}

	return nil
}

func (u *attachmentUsecase) DeleteAttachment(ID string) *apperror.AppError {
	return nil
}
