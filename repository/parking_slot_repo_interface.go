package repository

import (
	"ParkEase/data/request"
	"ParkEase/model"
)

type ParkingSlotRepository interface {
	Save(parkingLot model.ParkingSlots) (int64, error)
	GetAvailableSlotsByLot() ([]model.ParkingSlots, error)
	UpdateSlotAvailableStatus(req request.ParkVehicleRequest) error
	UpdateInMaintenanceStatus(req request.SlotMaintenanceRequest) error
}
