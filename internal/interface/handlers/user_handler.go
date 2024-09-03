package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/usecases"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(userUsecase usecases.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) InsertUser(c *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) UpdateUserByID(c *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) DeleteUserByID(c *fiber.Ctx) error {
	return nil
}
