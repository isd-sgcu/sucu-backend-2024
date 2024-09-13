package usecases

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/apperror"
)

type UserUsecase interface {
	// super-admin method
	GetAllUsers(req *dtos.UserDTO) (*[]dtos.UserDTO, *apperror.AppError)
	GetUserByID(req *dtos.UserDTO, userID string) (*dtos.UserDTO, *apperror.AppError)
	CreateUser(req *dtos.UserDTO, createUserDTO *dtos.CreateUserDTO) *apperror.AppError
	UpdateUserByID(req *dtos.UserDTO, userID string, updateUserDTO *dtos.UpdateUserDTO) *apperror.AppError
	DeleteUserByID(req *dtos.UserDTO, userID string) *apperror.AppError

	// admin method
	UpdateProfile(req *dtos.UserDTO, updateUserDTO *dtos.UpdateUserDTO) *apperror.AppError
}
