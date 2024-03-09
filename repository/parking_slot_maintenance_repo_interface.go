package repository

import (
	"ParkEase/model"

	"gorm.io/gorm"
)

type ParkingSlotMaintenanceRepository interface {
	Save(tx *gorm.DB, slotManitenance model.ParkingSlotMaintenances) (int64, error)
	UpdateMaintenanceCompleted(tx *gorm.DB, id int64) error
}
