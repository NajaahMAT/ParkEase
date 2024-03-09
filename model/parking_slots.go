package model

import "time"

type ParkingSlots struct {
	SlotID        int64       `gorm:"primary_key;auto_increment;comment:'Unique identifier for the Parking Slot'"`
	LotID         int64       `gorm:"type:int;not null;comment:'Parking LotID'"`
	SlotNo        int         `gorm:"type:int;not null"`
	InMaintenance bool        `gorm:"type:boolean;default:false"`
	IsAvailable   bool        `gorm:"type:boolean;default:true"`
	CreatedDate   time.Time   `gorm:"autoCreateTime"`
	UpdatedDate   time.Time   `gorm:"autoUpdateTime"`
	ParkingLots   ParkingLots `gorm:"foreignKey:LotID;references:LotID"`
}
