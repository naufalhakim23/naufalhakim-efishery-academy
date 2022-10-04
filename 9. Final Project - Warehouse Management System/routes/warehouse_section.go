package routes

import (
	"warehouse-management-system-eFishery/handler"

	"github.com/labstack/echo/v4"
)

func RoutesSection(e *echo.Echo, warehouseSectionHandler *handler.WarehouseSectionHandler) {
	e.POST("/api/warehouse/section", warehouseSectionHandler.CreateWarehouseSection)
	e.GET("/api/warehouse/section", warehouseSectionHandler.GetAllWarehouseSection)
	e.GET("/api/warehouse/section/:id", warehouseSectionHandler.GetWarehouseSectionByID)
	e.PUT("/api/warehouse/section/:id", warehouseSectionHandler.UpdateWarehouseSection)
	e.DELETE("/api/warehouse/section/:id", warehouseSectionHandler.DeleteWarehouseSection)
}
