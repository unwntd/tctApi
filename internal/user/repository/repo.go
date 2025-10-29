package repository

import (
	"tctApi/internal/user"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(m *user.User) error
	FindByEmail(email string) (*user.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(u *user.User) error {
	return r.db.Create(u).Error
}

func (r *userRepository) FindByEmail(email string) (*user.User, error) {
	var user user.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
