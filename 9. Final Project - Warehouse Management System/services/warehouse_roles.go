package services

import (
	entity "warehouse-management-system-eFishery/entity/warehouse"
	"warehouse-management-system-eFishery/repository"
)

type InterfaceWarehouseRolesService interface {
	CreateWarehouseRoles(warehouseRoles entity.WarehouseRoles) (entity.WarehouseRoles, error)
	GetAllWarehouseRoles() ([]entity.WarehouseRoles, error)
	GetWarehouseRolesByID(id int) (entity.WarehouseRoles, error)
	UpdateWarehouseRoles(id int, warehouseRoles entity.WarehouseRoles) (entity.WarehouseRoles, error)
	DeleteWarehouseRoles(id int) error
}

type WarehouseRolesService struct {
	warehouseRolesRepository repository.InterfaceWarehouseRolesRepository
}

func NewWarehouseRolesService(warehouseRolesRepository repository.InterfaceWarehouseRolesRepository) *WarehouseRolesService {
	return &WarehouseRolesService{warehouseRolesRepository}
}

func (service WarehouseRolesService) CreateWarehouseRoles(warehouseRoles entity.WarehouseRoles) (entity.WarehouseRoles, error) {
	warehouseRoles, err := service.warehouseRolesRepository.Store(warehouseRoles)
	if err != nil {
		return warehouseRoles, err
	}
	return warehouseRoles, nil
}

func (service WarehouseRolesService) GetAllWarehouseRoles() ([]entity.WarehouseRoles, error) {
	warehouseRoles, err := service.warehouseRolesRepository.FindAll()
	if err != nil {
		return warehouseRoles, err
	}
	return warehouseRoles, nil
}

func (service WarehouseRolesService) GetWarehouseRolesByID(id int) (entity.WarehouseRoles, error) {
	warehouseRoles, err := service.warehouseRolesRepository.FindByID(id)
	if err != nil {
		return warehouseRoles, err
	}
	return warehouseRoles, nil
}

func (service WarehouseRolesService) UpdateWarehouseRoles(id int, warehouseReq entity.WarehouseRoles) (entity.WarehouseRoles, error) {
	warehouseRoles, err := service.warehouseRolesRepository.FindByID(id)
	if err != nil {
		return entity.WarehouseRoles{}, err
	}
	wrs := entity.WarehouseRoles{
		ID:          warehouseRoles.ID,
		Role:        warehouseReq.Role,
		Description: warehouseReq.Description,
	}
	warehouseRoles, err = service.warehouseRolesRepository.Update(wrs)
	if err != nil {
		return entity.WarehouseRoles{}, err
	}
	return warehouseRoles, nil
}

func (service WarehouseRolesService) DeleteWarehouseRoles(id int) error {
	err := service.warehouseRolesRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
