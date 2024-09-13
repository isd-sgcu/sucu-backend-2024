package usecases

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/apperror"
)

type MiddlewareUsecase interface {
	VerifyToken(token string) (*string, *apperror.AppError)
	GetMe(userID string) (*dtos.UserDTO, *apperror.AppError)
}
