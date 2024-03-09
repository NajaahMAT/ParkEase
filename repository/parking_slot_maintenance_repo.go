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

func (d *ParkingSlotMaintenanceRepositoryImpl) Save(slotManitenance model.ParkingSlotMaintenances) (int64, error) {
	tx := d.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, err
	}

	if err := tx.Create(&slotManitenance).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return int64(slotManitenance.MaintenanceID), nil
}

func (t *ParkingSlotMaintenanceRepositoryImpl) UpdateMaintenanceCompleted(id int64) error {
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
	if err := tx.Model(&model.ParkingSlotMaintenances{}).Where("maintenance_id = ? ", id).Update("maintenance_end", time.Now()).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
