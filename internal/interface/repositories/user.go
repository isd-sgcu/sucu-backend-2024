package repositories

import (
	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
)

type UserRepository interface {
	FindAllUsers(limit int, offset int) (*[]entities.User, error)
	FindUserByID(ID string) (*entities.User, error)
	InsertUser(user *entities.User) error
	UpdateUserByID(ID string, updateMap interface{}) error
	DeleteUserByID(ID string) error
}
