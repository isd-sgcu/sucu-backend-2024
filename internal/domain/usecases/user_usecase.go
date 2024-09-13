package usecases

import (
	"errors"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/response"
	"github.com/isd-sgcu/sucu-backend-2024/utils"
	"github.com/isd-sgcu/sucu-backend-2024/utils/constant"
	"go.uber.org/zap"
	"gorm.io/gorm"
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

func (u *userUsecase) GetAllUsers(req *dtos.UserDTO) (*[]dtos.UserDTO, *response.AppError) {
	return nil, nil
}

func (u *userUsecase) GetUserByID(req *dtos.UserDTO, userID string) (*dtos.UserDTO, *response.AppError) {
	return nil, nil
}

func (u *userUsecase) CreateUser(req *dtos.UserDTO, createUserDTO *dtos.CreateUserDTO) *response.AppError {
	role, err := utils.GetRole(createUserDTO.Role)
	if err != nil {
		u.logger.Named("CreateUser").Error(constant.ErrInvalidRole, zap.String("role", createUserDTO.Role), zap.Error(err))
		return response.BadRequestError(constant.ErrInvalidRole)
	}

	existingUser, err := u.userRepository.FindUserByID(createUserDTO.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Named("CreateUser").Error(constant.ErrFindUserByID, zap.String("userID", createUserDTO.ID), zap.Error(err))
		return response.NotFoundError(constant.ErrFindUserByID)
	}

	if existingUser != nil {
		u.logger.Named("CreateUser").Error(constant.ErrUserAlreadyExists, zap.String("userID", createUserDTO.ID))
		return response.BadRequestError(constant.ErrUserAlreadyExists)
	}

	hashedPassword, err := utils.HashPassword(createUserDTO.Password)
	if err != nil {
		u.logger.Named("CreateUser").Error(constant.ErrHashPasswordFailed, zap.String("userID", createUserDTO.ID), zap.Error(err))
		return response.InternalServerError(constant.ErrHashPasswordFailed)
	}

	newUser := &entities.User{
		ID:        createUserDTO.ID,
		FirstName: createUserDTO.FirstName,
		LastName:  createUserDTO.LastName,
		Password:  hashedPassword,
		RoleID:    role,
	}

	if err := u.userRepository.InsertUser(newUser); err != nil {
		u.logger.Named("CreateUser").Error(constant.ErrInsertUserFailed, zap.String("userID", req.ID), zap.Error(err))
		return response.InternalServerError(constant.ErrInsertUserFailed)
	}

	u.logger.Named("CreateUser").Info("Success: ", zap.String("user_id", newUser.ID))
	return nil
}

func (u *userUsecase) UpdateUserByID(req *dtos.UserDTO, userID string, updateUserDTO *dtos.UpdateUserDTO) *response.AppError {
	return nil
}

func (u *userUsecase) DeleteUserByID(req *dtos.UserDTO, userID string) *response.AppError {
	return nil
}

// admin method

func (u *userUsecase) UpdateProfile(req *dtos.UserDTO, updateUserDTO *dtos.UpdateUserDTO) *response.AppError {
	return nil
}
