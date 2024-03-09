package repository

import "ParkEase/model"

type ParkingSlotMaintenanceRepository interface {
	Save(slotManitenance model.ParkingSlotMaintenances) (int64, error)
	UpdateMaintenanceCompleted(id int64) error
}
