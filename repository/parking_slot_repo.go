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

func (d *ParkingSlotRepositoryImpl) Save(tx *gorm.DB, parkingSlot model.ParkingSlots) (int64, error) {
	if err := tx.Create(&parkingSlot).Error; err != nil {
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

func (t *ParkingSlotRepositoryImpl) UpdateSlotAvailableStatus(tx *gorm.DB, req request.ParkVehicleRequest) error {
	// Proceed with updating the status
	if err := tx.Model(&model.ParkingSlots{}).Where("slot_id = ? ", req.SlotID).Update("is_available", req.IsSlotAvailable).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (t *ParkingSlotRepositoryImpl) UpdateInMaintenanceStatus(tx *gorm.DB, req request.SlotMaintenanceRequest) error {
	// Prepare a map of columns to update
	updates := map[string]interface{}{
		"in_maintenance": req.InMaintenance,
		"updated_date":   time.Now(),
	}

	// Proceed with updating the maintenance status
	if err := tx.Model(&model.ParkingSlots{}).Where("slot_id = ?", req.SlotID).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}
