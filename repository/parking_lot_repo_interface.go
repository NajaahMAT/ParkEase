package repository

import (
	"ParkEase/data/response"
	"ParkEase/model"

	"gorm.io/gorm"
)

type ParkingLotRepository interface {
	Save(tx *gorm.DB, parkingLot model.ParkingLots) (int64, error)
	GetParkingLotStatus() ([]response.ParkingLotStatusResponse, error)
}
