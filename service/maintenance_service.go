package service

import (
	"ParkEase/data/request"
	"ParkEase/model"
	"ParkEase/repository"
	"time"

	validator "github.com/go-playground/validator/v10"
)

type MaintenanceServiceImpl struct {
	ParkSlotRepository               repository.ParkingSlotRepository
	ParkingSlotMaintenanceRepository repository.ParkingSlotMaintenanceRepository
	Validate                         *validator.Validate
}

func NewMaintenanceServiceImpl(parkSlotRepository repository.ParkingSlotRepository, parkingSlotMaintenanceRepository repository.ParkingSlotMaintenanceRepository, validate *validator.Validate) MaintenanceService {
	return &MaintenanceServiceImpl{
		ParkSlotRepository:               parkSlotRepository,
		ParkingSlotMaintenanceRepository: parkingSlotMaintenanceRepository,
		Validate:                         validate,
	}
}

func (t *MaintenanceServiceImpl) PutSlotsIntoMaintenance(req request.SlotMaintenanceRequest) (maintenanceID int64, err error) {
	err = t.Validate.Struct(req)
	if err != nil {
		return
	}

	// Create new parking lot entry
	maintenanceModel := model.ParkingSlotMaintenances{
		SlotID:           req.SlotID,
		MaintenanceStart: time.Now(),
		Reason:           req.Reason,
	}

	maintenanceID, err = t.ParkingSlotMaintenanceRepository.Save(maintenanceModel)
	if err != nil {
		return 0, err
	}

	req.InMaintenance = true

	// update the slot available status to false in parking slots table
	err = t.ParkSlotRepository.UpdateInMaintenanceStatus(req)
	if err != nil {
		return 0, err
	}

	return maintenanceID, nil
}

func (t *MaintenanceServiceImpl) RestoreSlotsFromMaintenance(req request.SlotMaintenanceRequest, maintenanceID int64) (err error) {
	err = t.Validate.Struct(req)
	if err != nil {
		return err
	}

	//update the maintenace completed time
	err = t.ParkingSlotMaintenanceRepository.UpdateMaintenanceCompleted(maintenanceID)
	if err != nil {
		return err
	}

	// update the slot available status to false in parking slots table
	req.InMaintenance = false
	err = t.ParkSlotRepository.UpdateInMaintenanceStatus(req)
	if err != nil {
		return err
	}

	return nil
}
