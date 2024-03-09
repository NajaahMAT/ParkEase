package model

import "time"

type ParkingLots struct {
	LotID       int64     `gorm:"primary_key;auto_increment;comment:'Unique identifier for the ParkingLot'"`
	Name        string    `gorm:"type:varchar(50);not null;comment:'Name of the Parking Lot'"`
	TotalSlots  int       `gorm:"type:int;not null;comment:'Total slots in the parking lot'"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
}
