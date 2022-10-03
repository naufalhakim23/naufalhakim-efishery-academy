package handler

import (
	"net/http"
	response "warehouse-management-system-eFishery/entity/responseJson"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	"warehouse-management-system-eFishery/services"

	"github.com/labstack/echo/v4"
)

type WarehouseHandler struct {
	service *services.WarehouseService
}

func NewWarehouseHandler(service *services.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{service}
}

func (handler WarehouseHandler) CreateWarehouse(c echo.Context) error {
	var warehouse entity.Warehouse
	c.Bind(&warehouse)
	warehouse, err := handler.service.CreateWarehouse(warehouse)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create warehouse",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success to create warehouse",
		Data:    warehouse,
	})
}

func (handler WarehouseHandler) GetAllWarehouse(c echo.Context) error {
	warehouses, err := handler.service.GetAllWarehouse()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get list of e-Fishery Warehouses",
			Error:   err.Error(),
		})
	}
	if len(warehouses) == 0 {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get all list of e-Fishery Warehouses",
		Data:    warehouses,
	})
}
