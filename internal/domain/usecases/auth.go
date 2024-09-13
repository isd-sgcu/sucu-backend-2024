package usecases

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"
	"github.com/isd-sgcu/sucu-backend-2024/pkg/apperror"
)

type AuthUsecase interface {
	Login(loginUserDTO *dtos.LoginUserDTO) (*dtos.LoginResponseDTO, *apperror.AppError)
}
