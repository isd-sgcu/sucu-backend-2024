package handlers

import (
	"errors"
	"fmt"

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
	 getDocumentsDTO := dtos.GetDocumentsDTO{
		Page: c.QueryInt("page", 1),
		Limit: c.QueryInt("limit", 10),
		Query: c.Query("query"),
		Org: c.Query("org"),
		Type: c.Query("type"),
	}

	if getDocumentsDTO.Page < 1 {
		resp := response.NewResponseFactory(response.ERROR, errors.New("page need to be greater than 0"))
		resp.SendResponse(c, fiber.StatusBadRequest)
	}

	if getDocumentsDTO.Limit >= 20 || getDocumentsDTO.Limit < 0 {
		resp := response.NewResponseFactory(response.ERROR, errors.New("limit need to be between 0 and 20"))
		resp.SendResponse(c, fiber.StatusBadRequest)		
	}

	if getDocumentsDTO.Org != "sccu" && getDocumentsDTO.Org != "sgcu" && getDocumentsDTO.Org != "" {
		resp := response.NewResponseFactory(response.ERROR, errors.New("org need to be either sccu, sgcu or empty"))
		resp.SendResponse(c, fiber.StatusBadRequest)
	}

	if getDocumentsDTO.Type != "statistic" && getDocumentsDTO.Type != "budget" && getDocumentsDTO.Type != "announcement" && getDocumentsDTO.Type != "" {
		resp := response.NewResponseFactory(response.ERROR, errors.New("type need to be either statistic, budget, announcement or empty"))
		resp.SendResponse(c, fiber.StatusBadRequest)
	}

	documentsDTO, err := h.documentUsecase.GetAllDocuments(&getDocumentsDTO)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		resp.SendResponse(c, fiber.StatusInternalServerError)
	}

	paginationResponseDTO := dtos.PaginationResponse{
		Data: *documentsDTO,
		Page: fmt.Sprintf("%v", getDocumentsDTO.Page),
		Limit: fmt.Sprintf("%v", getDocumentsDTO.Limit),
		TotalPage: fmt.Sprintf("%v", len(*documentsDTO)),
	}

	resp := response.NewResponseFactory(response.SUCCESS, paginationResponseDTO)
	return resp.SendResponse(c, fiber.StatusAccepted)
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
	return nil
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
	return nil
}
