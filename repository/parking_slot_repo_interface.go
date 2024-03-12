package repository

import (
	"ParkEase/data/request"
	"ParkEase/model"

	"gorm.io/gorm"
)

type ParkingSlotRepository interface {
	Save(tx *gorm.DB, parkingLot model.ParkingSlots) (int64, error)
	GetAvailableSlotsByLot() ([]model.ParkingSlots, error)
	UpdateSlotAvailableStatus(tx *gorm.DB, req request.ParkVehicleRequest) error
	UpdateInMaintenanceStatus(tx *gorm.DB, req request.SlotMaintenanceRequest) error
	GetAvailableSlotsByCreteria(isSlotOdd bool) (*model.ParkingSlots, error)
}
