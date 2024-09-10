package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/usecases"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/response"
)

type AuthHandler struct {
	authUsecase usecases.AuthUsecase
}

func NewAuthHandler(authUsecase usecases.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

// Login godoc
// @Summary Log in user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param loginUserDTO body dtos.LoginUserDTO true "Login credentials"
// @Success 200 {object} response.Response{data=dtos.LoginResponseDTO}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var loginUserDTO dtos.LoginUserDTO
	if err := c.BodyParser(&loginUserDTO); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	loginResponseDTO, err := h.authUsecase.Login(&loginUserDTO)
	if err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusUnauthorized)
	}

	resp := response.NewResponseFactory(response.SUCCESS, loginResponseDTO)
	return resp.SendResponse(c, fiber.StatusOK)
}

// GetMe godoc
// @Summary Get current user profile
// @Tags Authentication
// @Produce json
// @Success 200 {object} response.Response{data=dtos.UserDTO}
// @Failure 500 {object} response.Response
// @Router /auth/me [get]
// @Security BearerAuth
func (h *AuthHandler) GetMe(c *fiber.Ctx) error {
	userDTO, ok := c.Locals("user").(*dtos.UserDTO)
	if !ok {
		resp := response.NewResponseFactory(response.ERROR, errors.New("not found user profile in context").Error())
		return resp.SendResponse(c, fiber.StatusInternalServerError)
	}

	resp := response.NewResponseFactory(response.SUCCESS, userDTO)
	return resp.SendResponse(c, fiber.StatusOK)
}
