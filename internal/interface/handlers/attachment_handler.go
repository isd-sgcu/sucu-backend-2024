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
