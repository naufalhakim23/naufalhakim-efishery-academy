package handler

import (
	"net/http"
	"strconv"
	response "warehouse-management-system-eFishery/entity/responseJson"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	service "warehouse-management-system-eFishery/services/warehouse"

	"github.com/labstack/echo/v4"
)

type WarehouseProductHandler struct {
	service *service.WarehouseProductService
}

func NewWarehouseProductHandler(service *service.WarehouseProductService) *WarehouseProductHandler {
	return &WarehouseProductHandler{service}
}

func (handler WarehouseProductHandler) CreateWarehouseProduct(c echo.Context) error {
	var warehouseProduct entity.WarehouseProducts
	c.Bind(&warehouseProduct)
	warehouseProduct, err := handler.service.CreateWarehouseProduct(warehouseProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create warehouse product",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success to create warehouse product",
		Data:    warehouseProduct,
	})
}

func (handler WarehouseProductHandler) GetAllWarehouseProduct(c echo.Context) error {
	warehouseProduct, err := handler.service.GetAllWarehouseProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get all warehouse product",
			Error:   err.Error(),
		})
	}
	if len(warehouseProduct) == 0 {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get all warehouse product",
		Data:    warehouseProduct,
	})
}

func (handler WarehouseProductHandler) GetWarehouseProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	warehouseProduct, err := handler.service.GetWarehouseProductByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get warehouse product by id",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get warehouse product by id",
		Data:    warehouseProduct,
	})
}

func (handler WarehouseProductHandler) GetWarehouseProductByPrice(c echo.Context) error {
	minPrice, _ := strconv.Atoi(c.Param("minPrice"))
	maxPrice, _ := strconv.Atoi(c.Param("maxPrice"))
	warehouseProduct, err := handler.service.GetWarehouseProductByPrice(minPrice, maxPrice)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get warehouse product by price",
			Error:   err.Error(),
		})
	}
	if len(warehouseProduct) == 0 {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	if minPrice > maxPrice {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Min price must be lower than max price",
		})
	}
	if minPrice < 0 || maxPrice < 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Price must be positive",
		})
	}
	if minPrice == 0 && maxPrice == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Price must be filled",
		})
	}
	if maxPrice == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Max price must be filled",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get warehouse product by price",
		Data:    warehouseProduct,
	})

}

func (handler WarehouseProductHandler) UpdateWarehouseProduct(c echo.Context) error {
	req := entity.UpdateWarehouseProducts{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update warehouse product",
			Error:   err.Error(),
		})
	}
	id, _ := strconv.Atoi(c.Param("id"))
	warehouseProduct, err := handler.service.UpdateWarehouseProduct(id, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update warehouse product",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to update warehouse product",
		Data:    warehouseProduct,
	})

}

func (handler WarehouseProductHandler) DeleteWarehouseProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := handler.service.DeleteWarehouseProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to delete warehouse product",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to delete warehouse product",
	})
}
