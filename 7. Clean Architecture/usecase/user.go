package usecase

import (
	"clean-architecture/entity"
	"clean-architecture/repository"
)

type IUserUsecase interface {
	CreateUser(user entity.UserRequest) (entity.User, error)
}
type UserUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{userRepository}
}
func (usecase UserUsecase) CreateUser(user entity.UserRequest) (entity.UserResponse, error) {
	u := entity.User{
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	users, err := usecase.userRepository.Store(u)
	if err != nil {
		return entity.UserResponse{}, err
	}
	userRes := entity.UserResponse{
		ID:       users.ID,
		Username: users.Username,
		Email:    users.Email,
		Phone:    users.Phone,
	}
	return userRes, nil
}
