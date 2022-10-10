package handler

import (
	"net/http"
	"strconv"
	response "warehouse-management-system-eFishery/entity/responseJson"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	services "warehouse-management-system-eFishery/services/warehouse"

	"github.com/labstack/echo/v4"
)

type WarehouseOrdersHandler struct {
	warehouseOrderService services.InterfaceWarehouseOrderService
}

func NewWarehouseOrdersHandler(warehouseOrderService services.InterfaceWarehouseOrderService) *WarehouseOrdersHandler {
	return &WarehouseOrdersHandler{warehouseOrderService}
}

func (handler WarehouseOrdersHandler) CreateWarehouseOrder(c echo.Context) error {
	var warehouseReq entity.CreateWarehouseOrders
	c.Bind(&warehouseReq)
	warehouseOrder, err := handler.warehouseOrderService.CreateWarehouseOrder(warehouseReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create warehouse order",
			Error:   err.Error(),
		})
	}
	if warehouseOrder.OrderId == 0 || warehouseOrder.ProductMark == "" || warehouseOrder.ProductStatus == "" || warehouseOrder.WorkerUUID == "" || warehouseOrder.WarehouseId == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create warehouse order",
			Error:   "Please fill all the fields",
		})
	}
	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success to create warehouse order",
		Data:    warehouseOrder,
	})
}

func (handler WarehouseOrdersHandler) GetAllWarehouseOrder(c echo.Context) error {
	warehouseOrders, err := handler.warehouseOrderService.FindAllWarehouseOrder()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get all warehouse order",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get all warehouse order",
		Data:    warehouseOrders,
	})
}

func (handler WarehouseOrdersHandler) GetWarehouseOrderById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	warehouseOrder, err := handler.warehouseOrderService.FindWarehouseOrderById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get warehouse order by id",
			Error:   err.Error(),
		})
	}
	warehouseOrd, _ := handler.warehouseOrderService.FindAllWarehouseOrder()
	if id > len(warehouseOrd) || id < 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update warehouse order by id",
			Error:   "Id not found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get warehouse order by id",
		Data:    warehouseOrder,
	})
}

func (handler WarehouseOrdersHandler) GetWarehouseOrderByProductStatus(c echo.Context) error {
	status := c.Param("product_status")
	warehouseOrders, err := handler.warehouseOrderService.FindWarehouseOrderByProductStatus(status)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get warehouse order by product status",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get warehouse order by product status",
		Data:    warehouseOrders,
	})
}

func (handler WarehouseOrdersHandler) GetWarehouseOrderByProductMark(c echo.Context) error {
	mark := c.Param("product_mark")
	warehouseOrders, err := handler.warehouseOrderService.FindWarehouseOrderByProductMark(mark)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get warehouse order by product mark",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get warehouse order by product mark",
		Data:    warehouseOrders,
	})
}

func (handler WarehouseOrdersHandler) UpdateWarehouseOrderById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var warehouseReq entity.UpdateWarehouseOrders
	c.Bind(&warehouseReq)
	warehouseOrder, err := handler.warehouseOrderService.UpdateWarehouseOrderById(id, warehouseReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update warehouse order by id",
			Error:   err.Error(),
		})
	}
	if warehouseOrder.OrderId == 0 || warehouseOrder.ProductMark == "" || warehouseOrder.ProductStatus == "" || warehouseOrder.WorkerUUID == "" || warehouseOrder.WarehouseId == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update warehouse order by id",
			Error:   "Please fill all the fields",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to update warehouse order by id",
		Data:    warehouseOrder,
	})
}

func (handler WarehouseOrdersHandler) DeleteWarehouseOrderById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := handler.warehouseOrderService.DeleteWarehouseOrderById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to delete warehouse order by id",
			Error:   err.Error(),
		})
	}
	warehouseOrd, _ := handler.warehouseOrderService.FindAllWarehouseOrder()
	if id > len(warehouseOrd) || id < 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update warehouse order by id",
		})
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to delete warehouse order by id",
	})

}
