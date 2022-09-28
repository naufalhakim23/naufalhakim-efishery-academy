package routes

import (
	"exercise-clean-architecture/handler"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, userHandler *handler.UserHandler) {
	e.POST("/users", userHandler.CreateUser)
	e.GET("/users", userHandler.GetAllUser)
	e.GET("/users/:uuid", userHandler.GetUserByUUID)
	// e.PUT("/users/:uuid", userHandler.UpdateDataByUUID)
	e.DELETE("/users/:uuid", userHandler.DeleteUserByUUID)
}
