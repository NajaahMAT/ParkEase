package model

import "time"

type ParkingSlotMaintenances struct {
	MaintenanceID    int64        `gorm:"primary_key;auto_increment;comment:'Unique identifier for the Parking Slot Maintenance'"`
	SlotID           int64        `gorm:"not null"`
	MaintenanceStart time.Time    `gorm:"not null"`
	MaintenanceEnd   time.Time    `gorm:"default:9999-12-31 23:59:59"`
	Reason           string       `gorm:"type:varchar(200);not null"`
	ParkingSlots     ParkingSlots `gorm:"foreignKey:SlotID;references:SlotID"`
}
