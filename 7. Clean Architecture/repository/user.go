package repository

import (
	"clean-architecture/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Store(user entity.User) (entity.User, error)
}
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}
func (r UserRepository) Store(user entity.User) (entity.User, error) {
	if err := r.db.Debug().Create(&user).Error; err != nil {
	}
	return user, nil
}
