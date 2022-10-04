package handler

import (
	"net/http"
	"strconv"
	response "warehouse-management-system-eFishery/entity/responseJson"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	"warehouse-management-system-eFishery/services"

	"github.com/labstack/echo/v4"
)

type WarehouseCategoryHandler struct {
	service *services.WarehouseCategoryService
}

func NewWarehouseCategoryHandler(service *services.WarehouseCategoryService) *WarehouseCategoryHandler {
	return &WarehouseCategoryHandler{service}
}

func (handler WarehouseCategoryHandler) CreateWarehouseCategory(c echo.Context) error {
	var warehouseCategory entity.WarehouseCategories
	c.Bind(&warehouseCategory)
	warehouseCategory, err := handler.service.CreateWarehouseCategory(warehouseCategory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create warehouse category",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success to create warehouse category",
		Data:    warehouseCategory,
	})
}

func (handler WarehouseCategoryHandler) GetAllWarehouseCategory(c echo.Context) error {
	warehouseCategory, err := handler.service.GetAllCategory()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get list of e-Fishery Warehouse Categories",
			Error:   err.Error(),
		})
	}
	if len(warehouseCategory) == 0 {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get all list of e-Fishery Warehouse Categories",
		Data:    warehouseCategory,
	})
}

func (handler WarehouseCategoryHandler) GetWarehouseCategoryByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	warehouseCategory, err := handler.service.GetCategoryByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get warehouse category by id",
			Error:   err.Error(),
		})
	}
	if warehouseCategory.ID == 0 {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get warehouse category by id",
		Data:    warehouseCategory,
	})
}

func (handler WarehouseCategoryHandler) UpdateWarehouseCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	warehouseCategory := entity.WarehouseCategories{}
	if err := c.Bind(&warehouseCategory); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update warehouse category",
			Error:   err.Error(),
		})
	}
	warehouseCategory, err := handler.service.UpdateWarehouseCategory(id, warehouseCategory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update warehouse category",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to update warehouse category",
		Data:    warehouseCategory,
	})
}

func (handler WarehouseCategoryHandler) DeleteWarehouseCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := handler.service.DeleteWarehouseCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to delete warehouse category",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to delete warehouse category",
	})
}
