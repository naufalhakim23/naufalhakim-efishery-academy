package routes

import (
	handler "warehouse-management-system-eFishery/handler/supplier"

	"github.com/labstack/echo/v4"
)

func RoutesSupplierAddress(e *echo.Echo, supplierAddressHandler *handler.SupplierAddressHandler) {
	e.POST("/api/supplier/address", supplierAddressHandler.CreateSupplierAddress)
	e.GET("/api/supplier/address", supplierAddressHandler.GetAllSupplierAddress)
	e.GET("/api/supplier/address/:id", supplierAddressHandler.GetSupplierAddressById)
	e.PUT("/api/supplier/address/:id", supplierAddressHandler.UpdateSupplierAddress)
	e.DELETE("/api/supplier/address/:id", supplierAddressHandler.DeleteSupplierAddress)
}
