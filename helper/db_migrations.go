package helper

import (
	"ParkEase/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	// Perform auto-migration for all models
	err := db.Table("parking_lots").AutoMigrate(&model.ParkingLots{})
	if err != nil {
		return err
	}

	err = db.Table("parking_slots").AutoMigrate(&model.ParkingSlots{})
	if err != nil {
		return err
	}

	// Set up foreign key constraints for parking_slots table
	db.Exec("ALTER TABLE parking_slots ADD CONSTRAINT fk_parking_lots FOREIGN KEY (lot_id) REFERENCES parking_lots(lot_id)")

	err = db.Table("parking_fees").AutoMigrate(&model.ParkingFees{})
	if err != nil {
		return err
	}

	// Set up foreign key constraints for parking_fee table
	db.Exec("ALTER TABLE parking_fees ADD CONSTRAINT fk_parking_slots FOREIGN KEY (slots_id) REFERENCES parking_slots(slot_id)")

	err = db.Table("parking_slot_maintenances").AutoMigrate(&model.ParkingSlotMaintenances{})
	if err != nil {
		return err
	}

	// Set up foreign key constraints for parking_slot_maintenance table
	db.Exec("ALTER TABLE parking_slot_maintenances ADD CONSTRAINT fk_parking_slots FOREIGN KEY (slots_id) REFERENCES parking_slots(slot_id)")

	return nil
}
