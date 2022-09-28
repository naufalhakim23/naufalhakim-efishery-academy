package handler

import (
	"exercise-clean-architecture/entity"
	"exercise-clean-architecture/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}

func (handler UserHandler) CreateUser(c echo.Context) error {
	req := entity.CreateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	user, err := handler.userUsecase.CreateUser(req)
	if err != nil {
		return err
	}
	return c.JSON(201, user)
}

func (hanlder UserHandler) GetAllUser(c echo.Context) error {
	users, err := hanlder.userUsecase.GetAllUser()
	if err != nil {
		return c.JSON(500, err)
	}
	if len(users) == 0 {
		return c.JSON(404, "No data found")
	}
	return c.JSON(200, users)
}
