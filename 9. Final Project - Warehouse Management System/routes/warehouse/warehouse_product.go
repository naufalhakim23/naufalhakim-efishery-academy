package routes

import (
	handler "warehouse-management-system-eFishery/handler/warehouse"

	"github.com/labstack/echo/v4"
)

func RoutesProducts(e *echo.Echo, warehouseProductHandler *handler.WarehouseProductHandler) {
	e.POST("/api/warehouse/product", warehouseProductHandler.CreateWarehouseProduct)
	e.GET("/api/warehouse/product", warehouseProductHandler.GetAllWarehouseProduct)
	e.GET("/api/warehouse/product/:id", warehouseProductHandler.GetWarehouseProductByID)
	e.GET("/api/warehouse/product/:minPrice/:maxPrice", warehouseProductHandler.GetWarehouseProductByPrice)
	e.PUT("/api/warehouse/product/:id", warehouseProductHandler.UpdateWarehouseProduct)
	e.DELETE("/api/warehouse/product/:id", warehouseProductHandler.DeleteWarehouseProduct)
}
