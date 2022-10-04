package handler

import (
	"net/http"
	"strconv"
	response "warehouse-management-system-eFishery/entity/responseJson"
	entity "warehouse-management-system-eFishery/entity/supplier"
	services "warehouse-management-system-eFishery/services/supplier"

	"github.com/labstack/echo/v4"
)

type SupplierHandler struct {
	service *services.SupplierService
}

func NewSupplierHandler(service *services.SupplierService) *SupplierHandler {
	return &SupplierHandler{service}
}

func (handler *SupplierHandler) CreateSupplier(c echo.Context) error {
	var supplier entity.Supplier
	c.Bind(&supplier)
	supplier, err := handler.service.CreateSupplier(supplier)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create supplier",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success to create supplier",
		Data:    supplier,
	})
}

func (handler *SupplierHandler) GetAllSupplier(c echo.Context) error {
	suppliers, err := handler.service.GetAllSupplier()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get list of e-Fishery Suppliers",
			Error:   err.Error(),
		})
	}
	if len(suppliers) == 0 {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get all list of e-Fishery Suppliers",
		Data:    suppliers,
	})
}

func (handler *SupplierHandler) GetSupplierByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	supplier, err := handler.service.GetSupplierByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get supplier by id",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get supplier by id",
		Data:    supplier,
	})
}

func (hanndler *SupplierHandler) UpdateSupplier(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var supplierReq entity.UpdateSupplier
	c.Bind(&supplierReq)
	supplier, err := hanndler.service.UpdateSupplier(id, supplierReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update supplier",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to update supplier",
		Data:    supplier,
	})
}

func (handler *SupplierHandler) DeleteSupplier(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := handler.service.DeleteSupplier(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to delete supplier",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to delete supplier",
	})
}
