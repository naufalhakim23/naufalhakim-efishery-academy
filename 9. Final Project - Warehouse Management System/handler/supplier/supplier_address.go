package handler

import (
	"net/http"
	"strconv"
	response "warehouse-management-system-eFishery/entity/responseJson"
	entity "warehouse-management-system-eFishery/entity/supplier"
	services "warehouse-management-system-eFishery/services/supplier"

	"github.com/labstack/echo/v4"
)

type SupplierAddressHandler struct {
	service *services.SupplierAddressService
}

func NewSupplierAddressHandler(service *services.SupplierAddressService) *SupplierAddressHandler {
	return &SupplierAddressHandler{service}
}

func (handler SupplierAddressHandler) CreateSupplierAddress(c echo.Context) error {
	var supplierAddress entity.SupplierAddress
	c.Bind(&supplierAddress)
	supplierAddress, err := handler.service.CreateSupplierAddress(supplierAddress)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create supplier address",
			Error:   err.Error(),
		})
	}
	if supplierAddress.ID == 0 || supplierAddress.SupplierID == 0 || supplierAddress.FullAddress == "" || supplierAddress.City == "" || supplierAddress.Province == "" || supplierAddress.SubDistrict == "" || supplierAddress.District == "" || supplierAddress.PostalCode == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create supplier address",
			Error:   "Please fill all the field",
		})
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success to create supplier address",
		Data:    supplierAddress,
	})
}

func (handler SupplierAddressHandler) GetAllSupplierAddress(c echo.Context) error {
	supplierAddresses, err := handler.service.GetAllSupplierAddress()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get list of supplier address",
			Error:   err.Error(),
		})
	}
	if len(supplierAddresses) == 0 {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "No data found",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get all list of supplier address",
		Data:    supplierAddresses,
	})
}

func (handler SupplierAddressHandler) GetSupplierAddressById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	supplierAddress, err := handler.service.GetSupplierAddressByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get supplier address by id",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to get supplier address by id",
		Data:    supplierAddress,
	})
}

func (handler SupplierAddressHandler) UpdateSupplierAddress(c echo.Context) error {
	req := entity.UpdateSupplierAddress{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update supplier address",
			Error:   err.Error(),
		})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	supplierAddress, err := handler.service.UpdateSupplierAddress(req, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update supplier address",
			Error:   err.Error(),
		})
	}
	if supplierAddress.ID == 0 || supplierAddress.SupplierID == 0 || supplierAddress.FullAddress == "" || supplierAddress.City == "" || supplierAddress.Province == "" || supplierAddress.SubDistrict == "" || supplierAddress.District == "" || supplierAddress.PostalCode == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to create supplier address",
			Error:   "Please fill all the field",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to update supplier address",
		Data:    supplierAddress,
	})
}

func (handler SupplierAddressHandler) DeleteSupplierAddress(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := handler.service.DeleteSupplierAddress(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to delete supplier address",
			Error:   err.Error(),
		})
	}

	if id == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to delete supplier",
		})
	}

	supplierAddress, _ := handler.service.GetAllSupplierAddress()
	if id > len(supplierAddress) || id < 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to delete supplier address",
			Error:   "Supplier address not found",
		})
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Success to delete supplier address",
	})
}
