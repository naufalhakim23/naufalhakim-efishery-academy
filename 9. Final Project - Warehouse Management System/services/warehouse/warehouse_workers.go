package services

import (
	"time"
	config "warehouse-management-system-eFishery/config/auth"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	repository "warehouse-management-system-eFishery/repository/warehouse"

	"github.com/google/uuid"
)

type InterfaceWarehouseWorkerService interface {
	CreateWarehouseWorkerWarehouseAuth(entity.CreateWarehouseWorkers, entity.CreateWarehouseAuth) (entity.WarehouseWorkersResponse, error)
	GetAllUser() ([]entity.WarehouseWorkersResponse, error)
	GetWarehouseWorkerByUUID(uuid string) (entity.WarehouseWorkersResponse, error)
	UpdateWarehouseWorker(uuid string, warehouseReq entity.CreateWarehouseWorkers) (entity.WarehouseWorkersResponse, error)
	DeleteWarehouseWorker(uuid string) error
}

type WarehouseWorkerService struct {
	warehouseWorkerRepository repository.InterfaceWarehouseWorkersRepository
}

func NewWarehouseWorkerService(warehouseWorkerRepository repository.InterfaceWarehouseWorkersRepository) *WarehouseWorkerService {
	return &WarehouseWorkerService{warehouseWorkerRepository}
}

func (service WarehouseWorkerService) CreateWarehouseWorker(warehouseReq entity.CreateWarehouseWorkers, warehouseAuthChan entity.CreateWarehouseAuth) (entity.WarehouseWorkersResponse, error) {
	newUUID := uuid.New().String()
	hashedPassword, _ := config.HashPassword(warehouseReq.Password)
	warehouseWorker := entity.WarehouseWorkers{
		UUID:        newUUID,
		FirstName:   warehouseReq.FirstName,
		LastName:    warehouseReq.LastName,
		Phone:       warehouseReq.Phone,
		Email:       warehouseReq.Email,
		Password:    hashedPassword,
		WarehouseId: warehouseReq.WarehouseId,
		RolesId:     warehouseReq.RolesId,
		CreatedAt:   time.Now().Format(time.RFC3339Nano),
	}
	wssAuth := entity.WarehouseAuth{
		UUID:      newUUID,
		Email:     warehouseReq.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now().Format(time.RFC3339Nano),
	}
	warehouseWorker, err := service.warehouseWorkerRepository.Store(warehouseWorker, wssAuth)

	wssResponse := entity.WarehouseWorkersResponse{
		UUID:        warehouseWorker.UUID,
		FirstName:   warehouseWorker.FirstName,
		LastName:    warehouseWorker.LastName,
		Phone:       warehouseWorker.Phone,
		Email:       warehouseWorker.Email,
		WarehouseId: warehouseWorker.WarehouseId,
		RolesId:     warehouseWorker.RolesId,
	}
	if err != nil {
		return wssResponse, err
	}
	return wssResponse, nil

}
func (service WarehouseWorkerService) GetAllUser() ([]entity.WarehouseWorkersResponse, error) {
	warehouseWorkers, err := service.warehouseWorkerRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var warehouseWorkersResponse []entity.WarehouseWorkersResponse
	for _, warehouseWorker := range warehouseWorkers {
		warehouseWorkerResponse := entity.WarehouseWorkersResponse{
			UUID:        warehouseWorker.UUID,
			FirstName:   warehouseWorker.FirstName,
			LastName:    warehouseWorker.LastName,
			Phone:       warehouseWorker.Phone,
			Email:       warehouseWorker.Email,
			WarehouseId: warehouseWorker.WarehouseId,
			RolesId:     warehouseWorker.RolesId,
		}
		warehouseWorkersResponse = append(warehouseWorkersResponse, warehouseWorkerResponse)
	}
	return warehouseWorkersResponse, nil
}

func (service WarehouseWorkerService) GetWarehouseWorkerByUUID(uuid string) (entity.WarehouseWorkersResponse, error) {
	warehouseWorker, err := service.warehouseWorkerRepository.FindByUUID(uuid)
	if err != nil {
		return entity.WarehouseWorkersResponse{}, err
	}
	warehouseWorkerResponse := entity.WarehouseWorkersResponse{
		UUID:        warehouseWorker.UUID,
		FirstName:   warehouseWorker.FirstName,
		LastName:    warehouseWorker.LastName,
		Phone:       warehouseWorker.Phone,
		Email:       warehouseWorker.Email,
		WarehouseId: warehouseWorker.WarehouseId,
		RolesId:     warehouseWorker.RolesId,
	}
	return warehouseWorkerResponse, nil
}

func (service WarehouseWorkerService) UpdateWarehouseWorker(uuid string, warehouseReq entity.CreateWarehouseWorkers) (entity.WarehouseWorkersResponse, error) {
	warehouseWorker := entity.WarehouseWorkers{
		UUID:        uuid,
		FirstName:   warehouseReq.FirstName,
		LastName:    warehouseReq.LastName,
		Phone:       warehouseReq.Phone,
		Email:       warehouseReq.Email,
		WarehouseId: warehouseReq.WarehouseId,
		RolesId:     warehouseReq.RolesId,
	}
	warehouseWorker, err := service.warehouseWorkerRepository.UpdateByUUID(uuid, warehouseWorker)
	if err != nil {
		return entity.WarehouseWorkersResponse{}, err
	}
	warehouseWorkerResponse := entity.WarehouseWorkersResponse{
		UUID:        warehouseWorker.UUID,
		FirstName:   warehouseWorker.FirstName,
		LastName:    warehouseWorker.LastName,
		Phone:       warehouseWorker.Phone,
		Email:       warehouseWorker.Email,
		WarehouseId: warehouseWorker.WarehouseId,
		RolesId:     warehouseWorker.RolesId,
	}
	return warehouseWorkerResponse, nil

}

func (service WarehouseWorkerService) DeleteWarehouseWorker(uuid string) error {
	err := service.warehouseWorkerRepository.DeleteByUUID(uuid)
	if err != nil {
		return err
	}
	return nil
}
