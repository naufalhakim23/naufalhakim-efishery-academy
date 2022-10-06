package handler

import (
	"net/http"
	"strconv"
	response "warehouse-management-system-eFishery/entity/responseJson"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	services "warehouse-management-system-eFishery/services/warehouse"

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

func (handler WarehouseHandler) GetWarehouseByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	warehouse, err := handler.service.GetWarehouseByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get warehouse by id",
			Error:   err.Error(),
		})
	}
	if warehouse.ID == 0 || warehouse.ID != id {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get warehouse by id",
		Data:    warehouse,
	})
}

func (handler WarehouseHandler) UpdateWarehouse(c echo.Context) error {
	req := entity.UpdateWarehouse{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Request Body",
			Error:   err.Error(),
		})
	}
	id, _ := strconv.Atoi(c.Param("id"))
	// Searching id bugged this will not be permanent solution need to figure out where is the error
	w, err := handler.service.GetWarehouseByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get warehouse by id",
			Error:   err.Error(),
		})
	}
	if w.ID == 0 || w.ID != id {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	warehouse, err := handler.service.UpdateWarehouse(req, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update warehouse",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to update warehouse",
		Data:    warehouse,
	})
}

func (handler WarehouseHandler) DeleteWarehouse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := handler.service.DeleteWarehouse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to delete warehouse",
			Error:   err.Error(),
		})
	}
	if id == 0 {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to delete warehouse",
	})
}
