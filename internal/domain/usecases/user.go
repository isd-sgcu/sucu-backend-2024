package usecases

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/response"
)

type UserUsecase interface {
	// super-admin method
	GetAllUsers(req *dtos.UserDTO) (*[]dtos.UserDTO, error)
	GetUserByID(req *dtos.UserDTO, userID string) (*dtos.UserDTO, error)
	CreateUser(req *dtos.UserDTO, createUserDTO *dtos.CreateUserDTO) *response.AppError
	UpdateUserByID(req *dtos.UserDTO, userID string, updateUserDTO *dtos.UpdateUserDTO) error
	DeleteUserByID(req *dtos.UserDTO, userID string) error

	// admin method
	UpdateProfile(req *dtos.UserDTO, updateUserDTO *dtos.UpdateUserDTO) error
}
