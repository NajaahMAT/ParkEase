package repository

import (
	"ParkEase/model"
	"time"

	"gorm.io/gorm"
)

type ParkingSlotMaintenanceRepositoryImpl struct {
	Db *gorm.DB
}

func NewParkingSlotMaintenanceRepositoryImpl(Db *gorm.DB) ParkingSlotMaintenanceRepository {
	return &ParkingSlotMaintenanceRepositoryImpl{Db: Db}
}

func (d *ParkingSlotMaintenanceRepositoryImpl) Save(tx *gorm.DB, slotManitenance model.ParkingSlotMaintenances) (int64, error) {
	if err := tx.Create(&slotManitenance).Error; err != nil {
		return 0, err
	}

	return int64(slotManitenance.MaintenanceID), nil
}

func (t *ParkingSlotMaintenanceRepositoryImpl) UpdateMaintenanceCompleted(tx *gorm.DB, id int64) error {
	// Proceed with updating the status
	if err := tx.Model(&model.ParkingSlotMaintenances{}).Where("maintenance_id = ? ", id).Update("maintenance_end", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
