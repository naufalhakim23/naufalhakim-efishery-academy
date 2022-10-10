package handler

import (
	"net/http"
	response "warehouse-management-system-eFishery/entity/responseJson"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	services "warehouse-management-system-eFishery/services/warehouse"

	"github.com/labstack/echo/v4"
)

type WarehouseWorkerHandler struct {
	service *services.WarehouseWorkerService
}

func NewWarehouseWorkerHandler(service *services.WarehouseWorkerService) *WarehouseWorkerHandler {
	return &WarehouseWorkerHandler{service}
}

func (handler WarehouseWorkerHandler) CreateWarehouseWorker(c echo.Context) error {
	var warehouseReq entity.CreateWarehouseWorkers
	var warehouseAuthChan entity.CreateWarehouseAuth
	c.Bind(&warehouseReq)
	c.Bind(&warehouseAuthChan)
	warehouseWorker, err := handler.service.CreateWarehouseWorker(warehouseReq, warehouseAuthChan)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create warehouse worker",
			Error:   err.Error(),
		})
	}
	if warehouseWorker.Email == "" || warehouseWorker.FirstName == "" || warehouseWorker.LastName == "" || warehouseWorker.Phone == "" || warehouseWorker.RolesId == 0 || warehouseWorker.WarehouseId == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create supplier address",
			Error:   "Please fill all the field",
		})
	}
	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success to create warehouse worker",
		Data:    warehouseWorker,
	})
}

func (handler WarehouseWorkerHandler) GetAllWarehouseWorker(c echo.Context) error {
	warehouseWorkers, err := handler.service.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get all warehouse worker",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get all warehouse worker",
		Data:    warehouseWorkers,
	})
}

func (handler WarehouseWorkerHandler) GetWarehouseWorkerByUUID(c echo.Context) error {
	uuid := c.Param("uuid")
	warehouseWorker, err := handler.service.GetWarehouseWorkerByUUID(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get warehouse worker",
			Error:   err.Error(),
		})
	}

	if uuid == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get warehouse worker",
			Error:   "Please fill all the field",
		})
	}
	if uuid != warehouseWorker.UUID {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get warehouse worker",
			Error:   "Worker not found",
		})
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get warehouse worker by uuid",
		Data:    warehouseWorker,
	})
}

func (handler WarehouseWorkerHandler) UpdateWarehouseWorker(c echo.Context) error {
	uuid := c.Param("uuid")
	var warehouseReq entity.CreateWarehouseWorkers
	c.Bind(&warehouseReq)
	warehouseWorker, err := handler.service.UpdateWarehouseWorker(uuid, warehouseReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update warehouse worker",
			Error:   err.Error(),
		})
	}
	if warehouseWorker.Email == "" || warehouseWorker.FirstName == "" || warehouseWorker.LastName == "" || warehouseWorker.Phone == "" || warehouseWorker.RolesId == 0 || warehouseWorker.WarehouseId == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create supplier address",
			Error:   "Please fill all the field",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to update warehouse worker",
		Data:    warehouseWorker,
	})
}

func (handler WarehouseWorkerHandler) DeleteWarehouseWorker(c echo.Context) error {
	uuid := c.Param("uuid")
	err := handler.service.DeleteWarehouseWorker(uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to delete warehouse worker",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to delete warehouse worker",
	})
}
