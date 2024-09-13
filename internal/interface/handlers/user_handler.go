package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/usecases"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/response"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(userUsecase usecases.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
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
	// omitted parse, validation

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
	return nil
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
	return nil
}
