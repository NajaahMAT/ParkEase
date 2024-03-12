package model

import "time"

type ParkingFees struct {
	ID            int64  `gorm:"primary_key;auto_increment"`
	SlotID        int64  `gorm:"not null"`
	VehicleNumber string `gorm:"type:varchar(50);not null"`
	//VehicleNumber    int          `gorm:"type:int;not null"`
	ParkingStartTime time.Time    `gorm:"not null"`
	ParkingEndTime   time.Time    `gorm:"default:9999-12-31 23:59:59"`
	ParkingFee       float64      `gorm:"type:decimal(10,6)"`
	ParkingSlots     ParkingSlots `gorm:"foreignKey:SlotID;references:SlotID"`
}
