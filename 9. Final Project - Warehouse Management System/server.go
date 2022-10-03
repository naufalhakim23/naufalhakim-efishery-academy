package main

import (
	config "warehouse-management-system-eFishery/config/database"
	"warehouse-management-system-eFishery/handler"
	"warehouse-management-system-eFishery/repository"
	"warehouse-management-system-eFishery/routes"
	"warehouse-management-system-eFishery/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.Connect()
	config.Migrate()
	e := echo.New()

	e.Use((middleware.Logger()))
	e.Use(middleware.Recover())

	warehouseRepository := repository.NewWarehouseRepository(config.DB)
	warehouseService := services.NewWarehouseService(warehouseRepository)
	warehouseHandler := handler.NewWarehouseHandler(warehouseService)
	routes.Routes(e, warehouseHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
