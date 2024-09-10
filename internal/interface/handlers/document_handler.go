package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/usecases"
)

type DocumentHandler struct {
	documentUsecase usecases.DocumentUsecase
}

func NewDocumentHandler(documentUsecase usecases.DocumentUsecase) *DocumentHandler {
	return &DocumentHandler{
		documentUsecase: documentUsecase,
	}
}

func (h *DocumentHandler) GetAllDocuments(c *fiber.Ctx) error {
	return nil
}

func (h *DocumentHandler) GetDocumentByID(c *fiber.Ctx) error {
	return nil
}

func (h *DocumentHandler) GetDocumentsByRole(c *fiber.Ctx) error {
	return nil
}

func (h *DocumentHandler) CreateDocument(c *fiber.Ctx) error {
	return nil
}

func (h *DocumentHandler) UpdateDocumentByID(c *fiber.Ctx) error {
	return nil
}

func (h *DocumentHandler) DeleteDocumentByID(c *fiber.Ctx) error {
	return nil
}
