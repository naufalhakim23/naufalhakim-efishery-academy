package handler

import (
	"net/http"
	"strconv"
	response "warehouse-management-system-eFishery/entity/responseJson"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	"warehouse-management-system-eFishery/services"

	"github.com/labstack/echo/v4"
)

type WarehouseAddressHandler struct {
	service *services.WarehouseAddressService
}

func NewWarehouseAddressHandler(service *services.WarehouseAddressService) *WarehouseAddressHandler {
	return &WarehouseAddressHandler{service}
}

func (handler WarehouseAddressHandler) CreateWarehouseAddress(c echo.Context) error {
	var warehouseAddress entity.WarehouseAddress
	c.Bind(&warehouseAddress)
	warehouseAddress, err := handler.service.CreateWarehouseAddress(warehouseAddress)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create warehouse address",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success to create warehouse address",
		Data:    warehouseAddress,
	})
}

func (handler WarehouseAddressHandler) GetAllWarehouseAddress(c echo.Context) error {
	warehouseAddresses, err := handler.service.GetAllAddress()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get list of e-Fishery Warehouse Addresses",
			Error:   err.Error(),
		})
	}
	if len(warehouseAddresses) == 0 {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get all list of e-Fishery Warehouse Addresses",
		Data:    warehouseAddresses,
	})
}

func (handler WarehouseAddressHandler) GetWarehouseAddressByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	warehouseAddress, err := handler.service.GetAddressByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get warehouse address by id",
			Error:   err.Error(),
		})
	}
	if warehouseAddress.ID == 0 {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get warehouse address by id",
		Data:    warehouseAddress,
	})
}

// Function UpdateWarehouseAddress still error
func (handler WarehouseAddressHandler) UpdateWarehouseAddress(c echo.Context) error {
	req := entity.UpdateWarehouseAddress{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update warehouse address",
			Error:   err.Error(),
		})
	}
	id, _ := strconv.Atoi(c.Param("id"))
	warehouseAddress, err := handler.service.UpdateWarehouseAddress(id, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update warehouse address",
			Error:   err.Error(),
		})
	}
	if warehouseAddress.ID == 0 {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to update warehouse address",
		Data:    warehouseAddress,
	})

}

func (handler WarehouseAddressHandler) DeleteWarehouseAddress(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := handler.service.DeleteWarehouseAddress(id)
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
