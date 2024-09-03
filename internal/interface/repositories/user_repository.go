package repositories

import (
	"context"

	"github.com/isd-sgcu/sucu-backend-2024/internal/domain/entities"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindAllUsers(ctx context.Context, roleID string) (*[]entities.User, error) {
	return nil, nil
}

func (r *userRepository) FindUserByID(ctx context.Context, ID string) (*entities.User, error) {
	var user entities.User

	if err := r.db.WithContext(ctx).First(&user, "id = ?", ID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) InsertUser(ctx context.Context, user *entities.User) error {
	return nil
}

func (r *userRepository) UpdateUserByID(ctx context.Context, ID string, updateMap interface{}) error {
	return nil
}

func (r *userRepository) DeleteUserByID(ctx context.Context, ID string) error {
	return nil
}
