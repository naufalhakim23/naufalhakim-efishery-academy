package routes

import (
	"warehouse-management-system-eFishery/handler"

	"github.com/labstack/echo/v4"
)

func RoutesAddress(e *echo.Echo, warehouseAddressHandler *handler.WarehouseAddressHandler) {
	e.POST("/api/warehouse/address", warehouseAddressHandler.CreateWarehouseAddress)
	e.GET("/api/warehouse/address", warehouseAddressHandler.GetAllWarehouseAddress)
	e.GET("/api/warehouse/address/:id", warehouseAddressHandler.GetWarehouseAddressByID)
	e.PUT("/api/warehouse/address/:id", warehouseAddressHandler.UpdateWarehouseAddress)
	e.DELETE("/api/warehouse/address/:id", warehouseAddressHandler.DeleteWarehouseAddress)

}
