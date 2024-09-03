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

func (h *AuthHandler) GetMe(c *fiber.Ctx) error {
	userDTO, ok := c.Locals("user").(*dtos.UserDTO)
	if !ok {
		resp := response.NewResponseFactory(response.ERROR, errors.New("not found user profile in context").Error())
		return resp.SendResponse(c, fiber.StatusInternalServerError)
	}

	resp := response.NewResponseFactory(response.SUCCESS, userDTO)
	return resp.SendResponse(c, fiber.StatusOK)
}
