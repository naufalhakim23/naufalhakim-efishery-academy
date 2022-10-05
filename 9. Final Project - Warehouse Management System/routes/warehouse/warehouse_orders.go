package routes

import (
	handler "warehouse-management-system-eFishery/handler/warehouse"

	"github.com/labstack/echo/v4"
)

func RouteOreders(e *echo.Echo, warehouseOrderHandler *handler.WarehouseOrdersHandler) {
	e.POST("/api/warehouse/order", warehouseOrderHandler.CreateWarehouseOrder)
	e.GET("/api/warehouse/order", warehouseOrderHandler.GetAllWarehouseOrder)
	e.GET("/api/warehouse/order/:id", warehouseOrderHandler.GetWarehouseOrderById)
	e.GET("/api/warehouse/order/:product_status", warehouseOrderHandler.GetWarehouseOrderByProductStatus)
	e.GET("/api/warehouse/order/:product_mark", warehouseOrderHandler.GetWarehouseOrderByProductMark)
	e.PUT("/api/warehouse/order/:id", warehouseOrderHandler.UpdateWarehouseOrderById)
	e.DELETE("/api/warehouse/order/:id", warehouseOrderHandler.DeleteWarehouseOrderById)
}
