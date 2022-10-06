package handler

import (
	"net/http"
	"strconv"
	response "warehouse-management-system-eFishery/entity/responseJson"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	services "warehouse-management-system-eFishery/services/warehouse"

	"github.com/labstack/echo/v4"
)

type WarehouseSectionHandler struct {
	service *services.WarehouseSectionService
}

func NewWarehouseSectionHandler(service *services.WarehouseSectionService) *WarehouseSectionHandler {
	return &WarehouseSectionHandler{service}
}

func (handler WarehouseSectionHandler) CreateWarehouseSection(c echo.Context) error {
	var warehouseSection entity.WarehouseSection
	c.Bind(&warehouseSection)
	warehouseSection, err := handler.service.CreateWarehouseSection(warehouseSection)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create warehouse section",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success to create warehouse section",
		Data:    warehouseSection,
	})
}

func (handler WarehouseSectionHandler) GetAllWarehouseSection(c echo.Context) error {
	warehouseSections, err := handler.service.GetAllWarehouseSection()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get list of e-Fishery Warehouse Sections",
			Error:   err.Error(),
		})
	}
	if len(warehouseSections) == 0 {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get all list of e-Fishery Warehouse Sections",
		Data:    warehouseSections,
	})
}

func (handler WarehouseSectionHandler) GetWarehouseSectionByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	warehouseSection, err := handler.service.GetWarehouseSectionByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get e-Fishery Warehouse Section",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get e-Fishery Warehouse Section",
		Data:    warehouseSection,
	})
}

func (handler WarehouseSectionHandler) UpdateWarehouseSection(c echo.Context) error {
	request := entity.UpdateWarehouseSection{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update e-Fishery Warehouse Section",
			Error:   err.Error(),
		})
	}
	id, _ := strconv.Atoi(c.Param("id"))
	warehouseSection, err := handler.service.UpdateWarehouseSection(id, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update e-Fishery Warehouse Section",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to update e-Fishery Warehouse Section",
		Data:    warehouseSection,
	})
}

func (handler WarehouseSectionHandler) DeleteWarehouseSection(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := handler.service.DeleteWarehouseSection(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to delete e-Fishery Warehouse Section",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to delete e-Fishery Warehouse Section",
	})
}
