package repository

import (
	"ParkEase/data/response"
	"ParkEase/model"
)

type ParkingLotRepository interface {
	Save(parkingLot model.ParkingLots) (int64, error)
	GetParkingLotStatus() ([]response.ParkingLotStatusResponse, error)
}
