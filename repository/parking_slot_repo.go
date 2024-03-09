package repository

import (
	"ParkEase/data/request"
	"ParkEase/model"
	"time"

	"gorm.io/gorm"
)

type ParkingSlotRepositoryImpl struct {
	Db *gorm.DB
}

func NewParkingSlotRepositoryImpl(Db *gorm.DB) ParkingSlotRepository {
	return &ParkingSlotRepositoryImpl{Db: Db}
}

func (d *ParkingSlotRepositoryImpl) Save(parkingSlot model.ParkingSlots) (int64, error) {
	tx := d.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, err
	}

	if err := tx.Create(&parkingSlot).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return parkingSlot.SlotID, nil
}

func (d *ParkingSlotRepositoryImpl) GetAvailableSlotsByLot() ([]model.ParkingSlots, error) {
	var parkingSlots []model.ParkingSlots
	if err := d.Db.Where("is_available = ? AND in_maintenance = ?", true, false).Order("lot_id, slot_no").Find(&parkingSlots).Error; err != nil {
		return nil, err
	}
	return parkingSlots, nil
}

func (t *ParkingSlotRepositoryImpl) UpdateSlotAvailableStatus(req request.ParkVehicleRequest) error {
	tx := t.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// Proceed with updating the status
	if err := tx.Model(&model.ParkingSlots{}).Where("slot_id = ? ", req.SlotID).Update("is_available", req.IsSlotAvailable).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (t *ParkingSlotRepositoryImpl) UpdateInMaintenanceStatus(req request.SlotMaintenanceRequest) error {
	tx := t.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// Prepare a map of columns to update
	updates := map[string]interface{}{
		"in_maintenance": req.InMaintenance,
		"updated_date":   time.Now(),
	}

	// Proceed with updating the maintenance status
	if err := tx.Model(&model.ParkingSlots{}).Where("slot_id = ?", req.SlotID).Updates(updates).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
