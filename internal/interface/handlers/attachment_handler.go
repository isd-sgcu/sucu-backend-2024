package handlers

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/usecases"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/response"
	"github.com/isd-sgcu/sucu-backend-2024/utils"
	"github.com/isd-sgcu/sucu-backend-2024/utils/constant"
)

type AttachmentHandler struct {
	attachmentUsecase usecases.AttachmentUsecase
}

func NewAttachmentHandler(attachmentUsecase usecases.AttachmentUsecase) *AttachmentHandler {
	return &AttachmentHandler{
		attachmentUsecase: attachmentUsecase,
	}
}

// GetAllAttachments godoc
// @Summary Get all attachments
// @Tags Attachments
// @Produce json
// @Success 200 {object} []dtos.AttachmentDTO
// @Failure 400 {object} response.Response
// @Router /attachments [get]
func (h *AttachmentHandler) GetAllAttachments(c *fiber.Ctx) error {
	getallAttachmentsDTO := dtos.GetAllAttachmentsDTO{
		Page:         	c.QueryInt("page", 1),
		PageSize:     	c.QueryInt("page_size", 10),
		DisplayName:  	c.Query("name"),
		AttachmentType: c.Query("attachment_type"),
	}

	var errors []string

	if !utils.ValidateDocType(getallAttachmentsDTO.AttachmentType) {
		errors = append(errors, constant.ErrInvalidAttachmentType)
	}

	if ps := getallAttachmentsDTO.PageSize; ps > constant.MAX_PAGE_SIZE || ps < 0 {
		errors = append(errors, constant.ErrInvalidPageSize)
	}

	startTime, err1 := time.Parse(time.RFC3339, c.Query("start_time", time.Time{}.UTC().Format(time.RFC3339)))
	endTime, err2 := time.Parse(time.RFC3339, c.Query("end_time", time.Now().UTC().Format(time.RFC3339)))
	if err1 != nil || err2 != nil {
		errors = append(errors, constant.ErrInvalidTimeFormat)
	}

	if len(errors) != 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errors, ", "))
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	getallAttachmentsDTO.StartTime = startTime
	getallAttachmentsDTO.EndTime = endTime

	paginationResp, err := h.attachmentUsecase.GetAllAttachments(&getallAttachmentsDTO)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, err.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, paginationResp)
	return resp.SendResponse(c, fiber.StatusOK)
}

// GetAllAttachmentsByRole godoc
// @Summary Get all attachments by role
// @Tags Attachments
// @Produce json
// @Param role_id path string true "Role of the user"
// @Success 200 {object} []dtos.AttachmentDTO
// @Failure 400 {object} response.Response
// @Router /attachments/role/{role_id} [get]
func (h *AttachmentHandler) GetAllAttachmentsByRole(c *fiber.Ctx) error {
	getallAttachmentsByRole := dtos.GetAllAttachmentsByRoleDTO{
		Page:         	c.QueryInt("page", 1),
		PageSize:     	c.QueryInt("page_size", 10),
		DisplayName:  	c.Query("name"),
		AttachmentType: c.Query("attachment_type"),
		Role:         	c.Params("role"),
	}

	var errors []string

	if !utils.ValidateDocType(getallAttachmentsByRole.AttachmentType) {
		errors = append(errors, constant.ErrInvalidAttachmentType)
	}

	if ps := getallAttachmentsByRole.PageSize; ps > constant.MAX_PAGE_SIZE || ps < 0 {
		errors = append(errors, constant.ErrInvalidPageSize)
	}

	if role := getallAttachmentsByRole.Role; role == "" || !utils.ValidateRole(role) {
		errors = append(errors, constant.ErrInvalidRole)
	}

	startTime, err1 := time.Parse(time.RFC3339, c.Query("start_time", time.Time{}.UTC().Format(time.RFC3339)))
	endTime, err2 := time.Parse(time.RFC3339, c.Query("end_time", time.Now().UTC().Format(time.RFC3339)))
	if err1 != nil || err2 != nil {
		errors = append(errors, constant.ErrInvalidTimeFormat)
	}

	if len(errors) != 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errors, ", "))
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	getallAttachmentsByRole.StartTime = startTime
	getallAttachmentsByRole.EndTime = endTime

	paginationResp, err := h.attachmentUsecase.GetAllAttachmentsByRole(&getallAttachmentsByRole)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, err.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, paginationResp)
	return resp.SendResponse(c, fiber.StatusOK)
}

// CreateAttachments godoc
// @Summary Create new attachments
// @Tags Attachments
// @Accept multipart/form-data
// @Produce json
// @Param document_id path string true "Document ID"
// @Param file formData file true "Attachment files"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /attachments/{document_id} [post]
func (h *AttachmentHandler) CreateAttachments(c *fiber.Ctx) error {
	documentID := strings.Trim(c.Params("document_id"), " ")

	form, err := c.MultipartForm()
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	if err := h.attachmentUsecase.CreateAttachments(documentID, form.File); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusInternalServerError)
	}

	resp := response.NewResponseFactory(response.SUCCESS, nil)
	return resp.SendResponse(c, fiber.StatusOK)
}

// DeleteAttachment godoc
// @Summary Delete an attachment by ID
// @Tags Attachments
// @Produce json
// @Param attachment_id path string true "Attachment ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /attachments/{attachment_id} [delete]
func (h *AttachmentHandler) DeleteAttachment(c *fiber.Ctx) error {
	return nil
}
