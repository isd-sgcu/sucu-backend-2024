package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/usecases"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/response"
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
	return nil
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
	return nil
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
	attachmentID := strings.Trim(c.Params("attachment_id"), " ")

	if err := h.attachmentUsecase.DeleteAttachment(attachmentID); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusInternalServerError)
	}

	resp := response.NewResponseFactory(response.SUCCESS, nil)
	return resp.SendResponse(c, fiber.StatusOK)
}
