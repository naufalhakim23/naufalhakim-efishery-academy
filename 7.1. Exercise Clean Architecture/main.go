package main

import (
	"exercise-clean-architecture/config"
	"exercise-clean-architecture/handler"
	"exercise-clean-architecture/repository"
	"exercise-clean-architecture/routes"
	"exercise-clean-architecture/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Connect()
	config.Migrate()

	e := echo.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	UserHandler := handler.NewUserHandler(userUsecase)

	routes.Routes(e, UserHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
