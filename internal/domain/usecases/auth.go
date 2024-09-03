package usecases

import "github.com/isd-sgcu/sucu-backend-2024/internal/interface/dtos"

type AuthUsecase interface {
	Login(loginUserDTO *dtos.LoginUserDTO) (*dtos.LoginResponseDTO, error)
}
