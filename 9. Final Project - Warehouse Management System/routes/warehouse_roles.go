package routes

import (
	"warehouse-management-system-eFishery/handler"

	"github.com/labstack/echo/v4"
)

func RoutesRoles(e *echo.Echo, warehouseRolesHandler *handler.WarehouseRolesHandler) {
	e.POST("/api/warehouse/roles", warehouseRolesHandler.CreateWarehouseRoles)
	e.GET("/api/warehouse/roles", warehouseRolesHandler.GetAllWarehouseRoles)
	e.GET("/api/warehouse/roles/:id", warehouseRolesHandler.GetWarehouseRolesByID)
	e.PUT("/api/warehouse/roles/:id", warehouseRolesHandler.UpdateWarehouseRoles)
	e.DELETE("/api/warehouse/roles/:id", warehouseRolesHandler.DeleteWarehouseRoles)
}
