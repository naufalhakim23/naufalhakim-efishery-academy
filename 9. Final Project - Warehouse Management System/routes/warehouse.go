package routes

import (
	"warehouse-management-system-eFishery/handler"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, warehouseHandler *handler.WarehouseHandler) {
	e.POST("/api/warehouse", warehouseHandler.CreateWarehouse)
	e.GET("/api/warehouse", warehouseHandler.GetAllWarehouse)
}
