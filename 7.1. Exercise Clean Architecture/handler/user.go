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
	req := entity.UserRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	user, err := handler.userUsecase.CreateUser(req)
	if err != nil {
		return err
	}
	return c.JSON(201, user)
}
