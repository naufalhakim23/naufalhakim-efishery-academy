package routes

import (
	handler "warehouse-management-system-eFishery/handler/warehouse"

	"github.com/labstack/echo/v4"
)

func RoutesWorkers(e *echo.Echo, warehouseWorkerHandler *handler.WarehouseWorkerHandler) {
	e.POST("/api/warehouse/worker", warehouseWorkerHandler.CreateWarehouseWorker)
	e.GET("/api/warehouse/worker", warehouseWorkerHandler.GetAllWarehouseWorker)
	e.GET("/api/warehouse/worker/:uuid", warehouseWorkerHandler.GetWarehouseWorkerByUUID)
	e.PUT("/api/warehouse/worker/:uuid", warehouseWorkerHandler.UpdateWarehouseWorker)
	e.DELETE("/api/warehouse/worker/:uuid", warehouseWorkerHandler.DeleteWarehouseWorker)
}
