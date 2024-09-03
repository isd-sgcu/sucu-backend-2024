package usecases

import "github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"

type MiddlewareUsecase interface {
	VerifyToken(token string) (*string, error)
	GetMe(userID string) (*dtos.UserDTO, error)
}
