package repositories

import (
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

func (r *userRepository) FindAllUsers(roleID string) (*[]entities.User, error) {
	return nil, nil
}

func (r *userRepository) FindUserByID(ID string) (*entities.User, error) {
	var user entities.User

	if err := r.db.First(&user, "id = ?", ID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) InsertUser(user *entities.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateUserByID(ID string, updateMap interface{}) error {
	return nil
}

func (r *userRepository) DeleteUserByID(ID string) error {
	return nil
}
