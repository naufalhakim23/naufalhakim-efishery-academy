package usecase

import (
	"exercise-clean-architecture/entity"
	"exercise-clean-architecture/repository"
)

type InterfaceUserUsercase interface {
	CreateUser(user entity.UserRequest) (entity.User, error)
	GetAllUser() ([]entity.User, error)
}

type UserUsecase struct {
	userRepository repository.InterfaceUserRepository
}

func NewUserUsecase(userRepository repository.InterfaceUserRepository) *UserUsecase {
	return &UserUsecase{userRepository}

}

func (usecase UserUsecase) CreateUser(user entity.UserRequest) (entity.UserResponse, error) {
	u := entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     user.Phone,
	}

	users, err := usecase.userRepository.Store(u)

	if err != nil {
		return entity.UserResponse{}, err
	}
	userRes := entity.UserResponse{
		ID:        users.ID,
		Username:  users.Username,
		Email:     users.Email,
		Phone:     users.Phone,
		FirstName: users.FirstName,
		LastName:  users.LastName,
	}

	return userRes, nil
}

func (usecase UserUsecase) GetAllUser() ([]entity.UserResponse, error) {
	users, err := usecase.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	userRes := []entity.UserResponse{}
	for _, user := range users {
		userRes = append(userRes, entity.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Phone:     user.Phone,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	}

	return userRes, nil

}
