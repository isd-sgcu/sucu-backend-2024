package usecases

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/apperror"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"
	"github.com/isd-sgcu/sucu-backend-2024/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	cfg            config.Config
	logger         *zap.Logger
	userRepository repositories.UserRepository
}

func NewAuthUsecase(cfg config.Config, logger *zap.Logger, userRepository repositories.UserRepository) AuthUsecase {
	return &authUsecase{
		cfg:            cfg,
		logger:         logger,
		userRepository: userRepository,
	}
}

func (u *authUsecase) Login(loginUserDTO *dtos.LoginUserDTO) (*dtos.LoginResponseDTO, *apperror.AppError) {
	existedUser, err := u.userRepository.FindUserByID(loginUserDTO.StudentID)
	if err != nil {
		u.logger.Named("Login").Error("Find user by ID: ", zap.Error(err))
		return nil, apperror.NotFoundError("this user is not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(loginUserDTO.Password)); err != nil {
		u.logger.Named("Login").Error("Compare hash and password: ", zap.Error(err))
		return nil, apperror.BadRequestError("this given password is not match with existed password")
	}

	tokenStr, err := utils.JwtSignAccessToken(existedUser.ID, u.cfg.GetJwt().AccessTokenSecret, u.cfg.GetJwt().AccessTokenExpiration)
	if err != nil {
		u.logger.Named("Login").Error("Jwt sign access token: ", zap.Error(err))
		return nil, apperror.InternalServerError("error while sign access token")
	}

	u.logger.Named("Login").Info("Success: ", zap.String("token", *tokenStr))
	return &dtos.LoginResponseDTO{AccessToken: *tokenStr}, nil
}
