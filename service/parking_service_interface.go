package service

import (
	"ParkEase/data/request"
	"ParkEase/data/response"
	"time"
)

type ParkingService interface {
	Create(req request.CreateParkingLotRequest) (int64, error)
	GetAvailableSlotsByLot() (resp response.AvailableSlotsResponse, err error)
	ParkVehicle(req request.ParkVehicleRequest) (feeID int64, err error)
	UnParkVehicle(id int64) (resp response.UnParkVehicleResponse, err error)
	GetParkingLotStatus() (resp []response.ParkingLotStatusResponse, err error)
	GetParkingStats(date time.Time) (resp response.ParkingStatsResponse, err error)
}
