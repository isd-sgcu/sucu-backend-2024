package usecases

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/response"
	"github.com/isd-sgcu/sucu-backend-2024/utils/constant"
	"go.uber.org/zap"
)

type userUsecase struct {
	cfg            config.Config
	logger         *zap.Logger
	userRepository repositories.UserRepository
}

func NewUserUsecase(cfg config.Config, logger *zap.Logger, userRepository repositories.UserRepository) UserUsecase {
	return &userUsecase{
		cfg:            cfg,
		logger:         logger,
		userRepository: userRepository,
	}
}

// super-admin method

func (u *userUsecase) GetAllUsers(req *dtos.UserDTO) (*[]dtos.UserDTO, error) {
	return nil, nil
}

func (u *userUsecase) GetUserByID(req *dtos.UserDTO, userID string) (*dtos.UserDTO, error) {
	return nil, nil
}

func (u *userUsecase) CreateUser(req *dtos.UserDTO, createUserDTO *dtos.CreateUserDTO) *response.AppError {
	return response.BadRequestError(constant.ErrInvalidRole)
}

func (u *userUsecase) UpdateUserByID(req *dtos.UserDTO, userID string, updateUserDTO *dtos.UpdateUserDTO) error {
	return nil
}

func (u *userUsecase) DeleteUserByID(req *dtos.UserDTO, userID string) error {
	return nil
}

// admin method

func (u *userUsecase) UpdateProfile(req *dtos.UserDTO, updateUserDTO *dtos.UpdateUserDTO) error {
	return nil
}
