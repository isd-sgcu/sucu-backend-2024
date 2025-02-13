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

func (r *userRepository) FindAllUsers(limit int, offset int) (*[]entities.User, error) {
	var users []entities.User

	if err := r.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *userRepository) FindUserByID(ID string) (*entities.User, error) {
	var user entities.User

	if err := r.db.First(&user, "id = ?", ID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) InsertUser(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) UpdateUserByID(ID string, updateMap interface{}) error {
	return r.db.Model(&entities.User{}).Where("id = ?", ID).Updates(updateMap).Error
}

func (r *userRepository) DeleteUserByID(ID string) error {
	return r.db.Where("id = ?", ID).Delete(&entities.User{}).Error
}
