package repositories

import (
	"github.com/isd-sgcu/sucu-backend-2024/pkg/config"

	"gorm.io/gorm"
)

type repository struct {
	UserRepository UserRepository
}

func NewRepository(db *gorm.DB, cfg config.Config) Repository {
	return &repository{
		UserRepository: NewUserRepository(db),
	}
}

func (r *repository) User() UserRepository {
	return r.UserRepository
}
