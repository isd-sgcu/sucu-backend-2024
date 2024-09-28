package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/usecases"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/response"
)

type DocumentHandler struct {
	documentUsecase usecases.DocumentUsecase
}

func NewDocumentHandler(documentUsecase usecases.DocumentUsecase) *DocumentHandler {
	return &DocumentHandler{
		documentUsecase: documentUsecase,
	}
}

// GetAllDocuments godoc
// @Summary Get all documents
// @Tags Documents
// @Produce json
// @Success 200 {object} response.Response{data=[]dtos.DocumentDTO}
// @Failure 500 {object} response.Response
// @Router /documents [get]
func (h *DocumentHandler) GetAllDocuments(c *fiber.Ctx) error {
	return nil
}

// GetDocumentByID godoc
// @Summary Get document by ID
// @Tags Documents
// @Produce json
// @Param document_id path string true "Document ID"
// @Success 200 {object} response.Response{data=dtos.DocumentDTO}
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /documents/{document_id} [get]
func (h *DocumentHandler) GetDocumentByID(c *fiber.Ctx) error {
	return nil
}

// GetDocumentsByRole godoc
// @Summary Get documents by user role
// @Tags Documents
// @Produce json
// @Param role_id path string true "User role"
// @Success 200 {object} response.Response{data=[]dtos.DocumentDTO}
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /documents/role/{role_id} [get]
func (h *DocumentHandler) GetDocumentsByRole(c *fiber.Ctx) error {
	return nil
}

// CreateDocument godoc
// @Summary Create a new document
// @Tags Documents
// @Accept json
// @Produce json
// @Param document body dtos.CreateDocumentDTO true "Document data"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /documents [post]
func (h *DocumentHandler) CreateDocument(c *fiber.Ctx) error {
	return nil
}

// UpdateDocumentByID godoc
// @Summary Update document by ID
// @Tags Documents
// @Accept json
// @Produce json
// @Param document_id path string true "Document ID"
// @Param document body dtos.UpdateDocumentDTO true "Updated document data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /documents/{document_id} [put]
func (h *DocumentHandler) UpdateDocumentByID(c *fiber.Ctx) error {
	documentID := c.Params("document_id")
	if documentID == "" {
		resp := response.NewResponseFactory(response.ERROR, "Document ID is required")
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	// Parse the request body into UpdateDocumentDTO
	var updateDocumentDTO dtos.UpdateDocumentDTO
	if err := c.BodyParser(&updateDocumentDTO); err != nil {
		resp := response.NewResponseFactory(response.ERROR, "Invalid request body")
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	err := h.documentUsecase.UpdateDocumentByID(documentID, updateDocumentDTO)
	if err != nil {
		if err.Error() == "document not found" {
			resp := response.NewResponseFactory(response.ERROR, "Document not found")
			return resp.SendResponse(c, fiber.StatusNotFound)
		}
		resp := response.NewResponseFactory(response.ERROR, "Failed to update document")
		return resp.SendResponse(c, fiber.StatusInternalServerError)
	}

	// Return success response
	resp := response.NewResponseFactory(response.SUCCESS, "Document updated successfully")
	return resp.SendResponse(c, fiber.StatusOK)
}

// DeleteDocumentByID godoc
// @Summary Delete document by ID
// @Tags Documents
// @Produce json
// @Param document_id path string true "Document ID"
// @Success 204 "No Content"
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /documents/{document_id} [delete]
func (h *DocumentHandler) DeleteDocumentByID(c *fiber.Ctx) error {
	documentID := c.Params("document_id")
	if documentID == "" {
		resp := response.NewResponseFactory(response.ERROR, "Document ID is required")
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	err := h.documentUsecase.DeleteDocumentByID(documentID)
	if err != nil {
		if err.Error() == "document not found" {
			resp := response.NewResponseFactory(response.ERROR, "Document not found")
			return resp.SendResponse(c, fiber.StatusNotFound)
		}
		resp := response.NewResponseFactory(response.ERROR, "Failed to delete document")
		return resp.SendResponse(c, fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
