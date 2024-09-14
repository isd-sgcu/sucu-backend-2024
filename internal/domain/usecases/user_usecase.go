package usecases

import (
	"errors"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/apperror"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"
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

func (u *userUsecase) GetAllUsers(req *dtos.UserDTO) (*[]dtos.UserDTO, *apperror.AppError) {
	return nil, nil
}

func (u *userUsecase) GetUserByID(req *dtos.UserDTO, userID string) (*dtos.UserDTO, *apperror.AppError) {
	return nil, nil
}

func (u *userUsecase) CreateUser(req *dtos.UserDTO, createUserDTO *dtos.CreateUserDTO) *apperror.AppError {
	role, err := utils.GetRole(req.Role)
	if err != nil {
		u.logger.Named("CreateUser").Error(constant.ErrInvalidRole, zap.String("role", createUserDTO.Role), zap.Error(err))
		return apperror.BadRequestError(constant.ErrInvalidRole)
	}

	existingUser, err := u.userRepository.FindUserByID(createUserDTO.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Named("CreateUser").Error(constant.ErrFindUserByID, zap.String("userID", createUserDTO.ID), zap.Error(err))
		return apperror.NotFoundError(constant.ErrFindUserByID)
	}

	if existingUser != nil {
		u.logger.Named("CreateUser").Error(constant.ErrUserAlreadyExists, zap.String("userID", createUserDTO.ID))
		return apperror.BadRequestError(constant.ErrUserAlreadyExists)
	}

	hashedPassword, err := utils.HashPassword(createUserDTO.Password)
	if err != nil {
		u.logger.Named("CreateUser").Error(constant.ErrHashPasswordFailed, zap.String("userID", createUserDTO.ID), zap.Error(err))
		return apperror.InternalServerError(constant.ErrHashPasswordFailed)
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
		return apperror.InternalServerError(constant.ErrInsertUserFailed)
	}

	u.logger.Named("CreateUser").Info("Success: ", zap.String("user_id", newUser.ID))
	return nil
}

func (u *userUsecase) UpdateUserByID(req *dtos.UserDTO, userID string, updateUserDTO *dtos.UpdateUserDTO) *apperror.AppError {
	return nil
}

func (u *userUsecase) DeleteUserByID(req *dtos.UserDTO, userID string) *apperror.AppError {
	return nil
}

// admin method

func (u *userUsecase) UpdateProfile(req *dtos.UserDTO, updateUserDTO *dtos.UpdateUserDTO) *apperror.AppError {
	updateFields := make(map[string]interface{})

	if updateUserDTO.FirstName != "" {
		updateFields["first_name"] = updateUserDTO.FirstName
	}

	if updateUserDTO.LastName != "" {
		updateFields["last_name"] = updateUserDTO.LastName
	}

	if updateUserDTO.Password != "" {
		hashedPassword, err := utils.HashPassword(updateUserDTO.Password)
		if err != nil {
			u.logger.Named("UpdateProfile").Error(constant.ErrHashPasswordFailed, zap.String("userID", req.ID), zap.Error(err))
			return apperror.InternalServerError(constant.ErrHashPasswordFailed)
		}
		updateFields["password"] = hashedPassword
	}

	if len(updateFields) == 0 {
		return apperror.BadRequestError("No fields to update")
	}

	err := u.userRepository.UpdateUserByID(req.ID, updateFields)
	if err != nil {
		u.logger.Named("UpdateProfile").Error(constant.ErrUpdateUserByID, zap.String("userID", req.ID), zap.Error(err))
		return apperror.BadRequestError(constant.ErrUpdateUserByID)
	}
	return nil
}
