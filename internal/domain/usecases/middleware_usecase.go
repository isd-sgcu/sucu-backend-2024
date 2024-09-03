package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
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

func (u *middlewareUsecase) VerifyToken(token string) (*string, error) {
	claim, err := utils.JwtParseToken(token, u.cfg.GetJwt().AccessTokenSecret)
	if err != nil {
		u.logger.Named("VerifyToken").Error("Parsing token: ", zap.Error(err))
		return nil, err
	}

	// get userId in token
	userID, ok := claim["sub"].(string)
	if !ok {
		u.logger.Named("VerifyToken").Error("Getting user_id from claim: ", zap.Error(errors.New("error while getting user_id from claim")))
		return nil, errors.New("user id not found in token")
	}

	u.logger.Named("VerifyToken").Info("Success: ", zap.String("user_id", userID))
	return &userID, nil
}

func (u *middlewareUsecase) GetMe(userID string) (*dtos.UserDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := u.userRepository.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &dtos.UserDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.RoleID,
	}, nil
}
