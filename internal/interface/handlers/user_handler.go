package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/usecases"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/response"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/validator"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
	validator   validator.DTOValidator
}

func NewUserHandler(userUsecase usecases.UserUsecase, validator validator.DTOValidator) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
		validator:   validator,
	}
}

// GetAllUsers godoc
// @Summary Get all users
// @Tags Users
// @Produce json
// @Success 200 {object} response.Response{data=[]dtos.UserDTO}
// @Failure 500 {object} response.Response
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	return nil
}

// GetUserByID godoc
// @Summary Get user by ID
// @Tags Users
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} response.Response{data=dtos.UserDTO}
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/{user_id} [get]
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	return nil
}

// CreateUser godoc
// @Summary Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dtos.CreateUserDTO true "User data"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var createUserDTO dtos.CreateUserDTO
	if err := c.BodyParser(&createUserDTO); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	if errs := h.validator.Validate(createUserDTO); len(errs) > 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", "))
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	req := c.Locals("user").(*dtos.UserDTO)
	apperr := h.userUsecase.CreateUser(req, &createUserDTO)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, createUserDTO)
	return resp.SendResponse(c, fiber.StatusCreated)
}

// UpdateUserByID godoc
// @Summary Update user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param user body dtos.UpdateUserDTO true "Updated user data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/{user_id} [put]
func (h *UserHandler) UpdateUserByID(c *fiber.Ctx) error {
	var updateUserDTO dtos.UpdateUserDTO
	if err := c.BodyParser(&updateUserDTO); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	if errs := h.validator.Validate(updateUserDTO); len(errs) > 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", "))
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	req := c.Locals("user").(*dtos.UserDTO)
	userID := c.Params("user_id")

	apperr := h.userUsecase.UpdateUserByID(req, userID, &updateUserDTO)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, updateUserDTO)
	return resp.SendResponse(c, fiber.StatusOK)
}

// DeleteUserByID godoc
// @Summary Delete user by ID
// @Tags Users
// @Produce json
// @Param user_id path string true "User ID"
// @Success 204 "No Content"
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/{user_id} [delete]
func (h *UserHandler) DeleteUserByID(c *fiber.Ctx) error {
	req := c.Locals("user").(*dtos.UserDTO)
	userID := c.Params("user_id")

	apperr := h.userUsecase.DeleteUserByID(req, userID)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dtos.UpdateUserDTO true "Updated user data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router / [patch]
func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	var updateProfile dtos.UpdateUserDTO
	if err := c.BodyParser(&updateProfile); err != nil {
		resp := response.NewResponseFactory(response.ERROR, err.Error())
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	if errs := h.validator.Validate(updateProfile); len(errs) > 0 {
		resp := response.NewResponseFactory(response.ERROR, strings.Join(errs, ", "))
		return resp.SendResponse(c, fiber.StatusBadRequest)
	}

	req := c.Locals("user").(*dtos.UserDTO)
	apperr := h.userUsecase.UpdateProfile(req, &updateProfile)
	if apperr != nil {
		resp := response.NewResponseFactory(response.ERROR, apperr.Error())
		return resp.SendResponse(c, apperr.HttpCode)
	}

	resp := response.NewResponseFactory(response.SUCCESS, updateProfile)
	return resp.SendResponse(c, fiber.StatusOK)
}
