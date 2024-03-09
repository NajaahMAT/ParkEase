package repository

import (
	"ParkEase/data/response"
	"ParkEase/model"
	"time"
)

type ParkingFeeRepository interface {
	Save(parkingFee model.ParkingFees) (int64, error)
	GetParkingFeeByID(id int64) (resp model.ParkingFees, err error)
	UpdateParkingFees(req model.ParkingFees) error
	GetParkingStats(date time.Time) (resp response.ParkingStatsResponse, err error)
}
