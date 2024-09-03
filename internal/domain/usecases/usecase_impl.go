package usecases

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/repositories"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"

	"go.uber.org/zap"
)

type usecase struct {
	MiddlewareUsecase MiddlewareUsecase
	AuthUsecase       AuthUsecase
	UserUsecase       UserUsecase
}

func NewService(repo repositories.Repository, cfg config.Config, logger *zap.Logger) Usecase {
	return &usecase{
		MiddlewareUsecase: NewMiddlewareUsecase(cfg, logger.Named("MiddlewareSvc"), repo.User()),
		AuthUsecase:       NewAuthUsecase(cfg, logger.Named("AuthSvc"), repo.User()),
		UserUsecase:       NewUserUsecase(repo.User()),
	}
}

func (u *usecase) Middleware() MiddlewareUsecase {
	return u.MiddlewareUsecase
}

func (u *usecase) Auth() AuthUsecase {
	return u.AuthUsecase
}

func (u *usecase) User() UserUsecase {
	return u.UserUsecase
}
