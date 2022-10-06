package routes

import (
	handler "warehouse-management-system-eFishery/handler/warehouse"

	"github.com/labstack/echo/v4"
)

func RoutesCategory(e *echo.Echo, warehouseCategoryHandler *handler.WarehouseCategoryHandler) {
	e.POST("/api/warehouse/category", warehouseCategoryHandler.CreateWarehouseCategory)
	e.GET("/api/warehouse/category", warehouseCategoryHandler.GetAllWarehouseCategory)
	e.GET("/api/warehouse/category/:id", warehouseCategoryHandler.GetWarehouseCategoryByID)
	e.PUT("/api/warehouse/category/:id", warehouseCategoryHandler.UpdateWarehouseCategory)
	e.DELETE("/api/warehouse/category/:id", warehouseCategoryHandler.DeleteWarehouseCategory)
}
