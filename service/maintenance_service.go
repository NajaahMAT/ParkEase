package service

import (
	"ParkEase/data/request"
	"ParkEase/model"
	"ParkEase/repository"
	"time"

	validator "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type MaintenanceServiceImpl struct {
	ParkSlotRepository               repository.ParkingSlotRepository
	ParkingSlotMaintenanceRepository repository.ParkingSlotMaintenanceRepository
	Validate                         *validator.Validate
	Db                               *gorm.DB // Add Db field
}

func NewMaintenanceServiceImpl(db *gorm.DB, parkSlotRepository repository.ParkingSlotRepository, parkingSlotMaintenanceRepository repository.ParkingSlotMaintenanceRepository, validate *validator.Validate) MaintenanceService {
	return &MaintenanceServiceImpl{
		ParkSlotRepository:               parkSlotRepository,
		ParkingSlotMaintenanceRepository: parkingSlotMaintenanceRepository,
		Validate:                         validate,
		Db:                               db, // Assign the Db instance
	}
}

func (t *MaintenanceServiceImpl) PutSlotsIntoMaintenance(req request.SlotMaintenanceRequest) (maintenanceID int64, err error) {
	err = t.Validate.Struct(req)
	if err != nil {
		return
	}

	// Start a transaction
	tx := t.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create new parking lot entry
	maintenanceModel := model.ParkingSlotMaintenances{
		SlotID:           req.SlotID,
		MaintenanceStart: time.Now(),
		Reason:           req.Reason,
	}

	maintenanceID, err = t.ParkingSlotMaintenanceRepository.Save(tx, maintenanceModel)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	req.InMaintenance = true

	// update the slot available status to false in parking slots table
	err = t.ParkSlotRepository.UpdateInMaintenanceStatus(tx, req)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return maintenanceID, nil
}

func (t *MaintenanceServiceImpl) RestoreSlotsFromMaintenance(req request.SlotMaintenanceRequest, maintenanceID int64) (err error) {
	err = t.Validate.Struct(req)
	if err != nil {
		return err
	}

	// Start a transaction
	tx := t.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//update the maintenace completed time
	err = t.ParkingSlotMaintenanceRepository.UpdateMaintenanceCompleted(tx, maintenanceID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// update the slot available status to false in parking slots table
	req.InMaintenance = false
	err = t.ParkSlotRepository.UpdateInMaintenanceStatus(tx, req)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
