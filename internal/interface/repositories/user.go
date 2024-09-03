package repositories

import (
	"context"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
)

type UserRepository interface {
	FindAllUsers(ctx context.Context, roleID string) (*[]entities.User, error)
	FindUserByID(ctx context.Context, ID string) (*entities.User, error)
	InsertUser(ctx context.Context, user *entities.User) error
	UpdateUserByID(ctx context.Context, ID string, updateMap interface{}) error
	DeleteUserByID(ctx context.Context, ID string) error
}
