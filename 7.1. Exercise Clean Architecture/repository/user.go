package repository

import (
	"exercise-clean-architecture/entity"

	"gorm.io/gorm"
)

type InterfaceUserRepository interface {
	Store(user entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
	FindByUUID(uuid string) (entity.User, error)
}
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) InterfaceUserRepository {
	return &UserRepository{db}
}

// Storing data to database
func (r UserRepository) Store(user entity.User) (entity.User, error) {
	if err := r.db.Debug().Create(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// Read data from database
func (r UserRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Debug().Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Find user using uuid from database
func (r UserRepository) FindByUUID(uuid string) (entity.User, error) {
	var user entity.User
	if err := r.db.Debug().Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}
