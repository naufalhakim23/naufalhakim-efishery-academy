package handler

import (
	"net/http"
	"strconv"
	response "warehouse-management-system-eFishery/entity/responseJson"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	services "warehouse-management-system-eFishery/services/warehouse"

	"github.com/labstack/echo/v4"
)

type WarehouseRolesHandler struct {
	service *services.WarehouseRolesService
}

func NewWarehouseRolesHandler(service *services.WarehouseRolesService) *WarehouseRolesHandler {
	return &WarehouseRolesHandler{service}
}

func (handler WarehouseRolesHandler) CreateWarehouseRoles(c echo.Context) error {
	var warehouseRoles entity.WarehouseRoles
	c.Bind(&warehouseRoles)
	warehouseRoles, err := handler.service.CreateWarehouseRoles(warehouseRoles)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create warehouse roles",
			Error:   err.Error(),
		})
	}
	if warehouseRoles.Role == "" || warehouseRoles.Description == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create warehouse roles",
			Error:   "Role or Description cannot be empty",
		})
	}
	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success to create warehouse roles",
		Data:    warehouseRoles,
	})
}

func (handler WarehouseRolesHandler) GetAllWarehouseRoles(c echo.Context) error {
	warehouseRoles, err := handler.service.GetAllWarehouseRoles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get list of e-Fishery Warehouse Roles",
			Error:   err.Error(),
		})
	}
	if len(warehouseRoles) == 0 {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get all list of e-Fishery Warehouse Roles",
		Data:    warehouseRoles,
	})
}

func (handler WarehouseRolesHandler) GetWarehouseRolesByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	warehouseRoles, err := handler.service.GetWarehouseRolesByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get warehouse roles by id",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get warehouse roles by id",
		Data:    warehouseRoles,
	})
}
func (handler WarehouseRolesHandler) UpdateWarehouseRoles(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	warehouseRoles := entity.WarehouseRoles{}
	if err := c.Bind(&warehouseRoles); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update warehouse roles",
			Error:   err.Error(),
		})
	}
	warehouseRoles, err := handler.service.UpdateWarehouseRoles(id, warehouseRoles)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update warehouse roles",
			Error:   err.Error(),
		})
	}
	if warehouseRoles.Role == "" || warehouseRoles.Description == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create warehouse roles",
			Error:   "Role or Description cannot be empty",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to update warehouse roles",
		Data:    warehouseRoles,
	})
}

func (handler WarehouseRolesHandler) DeleteWarehouseRoles(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := handler.service.DeleteWarehouseRoles(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete warehouse roles",
			Error:   err.Error(),
		})
	}
	warehouseRole, _ := handler.service.GetAllWarehouseRoles()
	if id > len(warehouseRole) || id < 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to delete warehouse roles",
			Error:   "ID not found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to delete warehouse roles",
	})
}
