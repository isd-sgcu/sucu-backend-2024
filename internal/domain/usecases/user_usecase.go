package usecases

import (
	"errors"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
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

func (u *userUsecase) GetAllUsers(req *dtos.UserDTO) (*[]dtos.UserDTO, error) {
	return nil, nil
}

func (u *userUsecase) GetUserByID(req *dtos.UserDTO, userID string) (*dtos.UserDTO, error) {
	return nil, nil
}

func (u *userUsecase) CreateUser(req *dtos.UserDTO, createUserDTO *dtos.CreateUserDTO) error {
	var role string
	switch req.Role {
	case constant.SGCU_SUPERADMIN:
		role = constant.SGCU_ADMIN
	case constant.SCCU_SUPERADMIN:
		role = constant.SCCU_ADMIN
	default:
		u.logger.Named("CreateUser").Error(constant.ErrInvalidRole.Error(), zap.String("role", req.Role))
		return constant.ErrInvalidRole
	}

	existingUser, err := u.userRepository.FindUserByID(createUserDTO.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Named("CreateUser").Error(constant.ErrFindUserByID.Error(), zap.String("userID", createUserDTO.ID), zap.Error(err))
		return constant.ErrFindUserByID
	}

	if existingUser != nil {
		u.logger.Named("CreateUser").Error(constant.ErrUserAlreadyExists.Error(), zap.String("userID", createUserDTO.ID))
		return constant.ErrUserAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(createUserDTO.Password)
	if err != nil {
		u.logger.Named("CreateUser").Error(constant.ErrHashPasswordFailed.Error(), zap.String("userID", createUserDTO.ID), zap.Error(err))
		return constant.ErrHashPasswordFailed
	}

	newUser := &entities.User{
		ID:        createUserDTO.ID,
		FirstName: createUserDTO.FirstName,
		LastName:  createUserDTO.LastName,
		Password:  hashedPassword,
		RoleID:    role,
	}

	if err := u.userRepository.InsertUser(newUser); err != nil {
		u.logger.Named("CreateUser").Error(constant.ErrInsertUserFailed.Error(), zap.String("userID", req.ID), zap.Error(err))
		return constant.ErrInsertUserFailed
	}

	u.logger.Named("CreateUser").Info("Success: ", zap.String("user_id", newUser.ID))
	return nil
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
