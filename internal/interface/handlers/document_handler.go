package handlers

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/usecases"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/response"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/validator"
)

type DocumentHandler struct {
	documentUsecase usecases.DocumentUsecase
	validator       validator.DTOValidator
}

func NewDocumentHandler(documentUsecase usecases.DocumentUsecase, validator validator.DTOValidator) *DocumentHandler {
	return &DocumentHandler{
		documentUsecase: documentUsecase,
		validator:       validator,
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
	getallDocumentsDTO := dtos.GetAllDocumentsDTO{
		Page:         c.QueryInt("page", 1),
		PageSize:     c.QueryInt("page_size", 10),
		Query:        c.Query("query"),
		Organization: c.Query("organization"),
		DocumentType: c.Query("document_type"),
		StartTime:    c.Query("Start_time", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).String()),
		EndTime:      c.Query("end_time", time.Now().String()),
	}

	paginationResp, err := h.documentUsecase.GetAllDocuments(&getallDocumentsDTO)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, err.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, paginationResp)
	return resp.SendResponse(c, fiber.StatusCreated)
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
	var CreateDocumentDTO dtos.CreateDocumentDTO
	if err := c.BodyParser(&CreateDocumentDTO); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	user := c.Locals("user").(*dtos.UserDTO)
	CreateDocumentDTO.UserID = user.ID

	if errs := h.validator.Validate(CreateDocumentDTO); len(errs) > 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", "))
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	apperr := h.documentUsecase.CreateDocument(&CreateDocumentDTO)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, "Document created successfully")
	return resp.SendResponse(c, fiber.StatusCreated)
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
