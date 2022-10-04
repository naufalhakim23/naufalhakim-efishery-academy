package routes

import (
	handler "warehouse-management-system-eFishery/handler/supplier"

	"github.com/labstack/echo/v4"
)

func RoutesSupplier(e *echo.Echo, supplierHandler *handler.SupplierHandler) {
	e.POST("/api/supplier", supplierHandler.CreateSupplier)
	e.GET("/api/supplier", supplierHandler.GetAllSupplier)
	e.GET("/api/supplier/:id", supplierHandler.GetSupplierByID)
	e.PUT("/api/supplier/:id", supplierHandler.UpdateSupplier)
	e.DELETE("/api/supplier/:id", supplierHandler.DeleteSupplier)
}
