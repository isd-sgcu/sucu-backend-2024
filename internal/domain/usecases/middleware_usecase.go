package usecases

import (
	"errors"

	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/apperror"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"
	"github.com/isd-sgcu/sucu-backend-2024/utils"
	"go.uber.org/zap"
)

type middlewareUsecase struct {
	cfg            config.Config
	logger         *zap.Logger
	userRepository repositories.UserRepository
}

func NewMiddlewareUsecase(cfg config.Config, logger *zap.Logger, userRepository repositories.UserRepository) MiddlewareUsecase {
	return &middlewareUsecase{
		cfg:            cfg,
		logger:         logger,
		userRepository: userRepository,
	}
}

func (u *middlewareUsecase) VerifyToken(token string) (*string, *apperror.AppError) {
	claim, err := utils.JwtParseToken(token, u.cfg.GetJwt().AccessTokenSecret)
	if err != nil {
		u.logger.Named("VerifyToken").Error("Parsing token: ", zap.Error(err))
		return nil, apperror.UnauthorizedError("invalid token")
	}

	// get userId in token
	userID, ok := claim["sub"].(string)
	if !ok {
		u.logger.Named("VerifyToken").Error("Getting user_id from claim: ", zap.Error(errors.New("error while getting user_id from claim")))
		return nil, apperror.InternalServerError("user id not found in token")
	}

	u.logger.Named("VerifyToken").Info("Success: ", zap.String("user_id", userID))
	return &userID, nil
}

func (u *middlewareUsecase) GetMe(userID string) (*dtos.UserDTO, *apperror.AppError) {
	user, err := u.userRepository.FindUserByID(userID)
	if err != nil {
		return nil, apperror.NotFoundError("user not found")
	}

	return &dtos.UserDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.RoleID,
	}, nil
}
