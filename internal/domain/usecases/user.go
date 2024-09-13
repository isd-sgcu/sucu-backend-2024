package usecases

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/response"
)

type UserUsecase interface {
	// super-admin method
	GetAllUsers(req *dtos.UserDTO) (*[]dtos.UserDTO, *response.AppError)
	GetUserByID(req *dtos.UserDTO, userID string) (*dtos.UserDTO, *response.AppError)
	CreateUser(req *dtos.UserDTO, createUserDTO *dtos.CreateUserDTO) *response.AppError
	UpdateUserByID(req *dtos.UserDTO, userID string, updateUserDTO *dtos.UpdateUserDTO) *response.AppError
	DeleteUserByID(req *dtos.UserDTO, userID string) *response.AppError

	// admin method
	UpdateProfile(req *dtos.UserDTO, updateUserDTO *dtos.UpdateUserDTO) *response.AppError
}
