package main

import (
	"clean-architecture/config"
	"clean-architecture/handler"
	"clean-architecture/repository"
	"clean-architecture/routes"
	"clean-architecture/usecase"

	"github.com/labstack/echo/v4"
)

// what the actual fuck is this.
func main() {
	config.Database()
	config.Migrate()

	e := echo.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	UserHandler := handler.NewUserHandler(userUsecase)

	routes.Routes(e, UserHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
