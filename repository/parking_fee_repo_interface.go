package repository

import (
	"ParkEase/data/response"
	"ParkEase/model"
	"time"

	"gorm.io/gorm"
)

type ParkingFeeRepository interface {
	Save(tx *gorm.DB, parkingFee model.ParkingFees) (int64, error)
	GetParkingFeeByID(id int64) (resp model.ParkingFees, err error)
	UpdateParkingFees(tx *gorm.DB, req model.ParkingFees) error
	GetParkingStats(date time.Time) (resp response.ParkingStatsResponse, err error)
}
