package repository

import (
	"ParkEase/data"
	"ParkEase/data/response"
	"ParkEase/model"
	"database/sql"

	"gorm.io/gorm"
)

type ParkingLotRepositoryImpl struct {
	Db *gorm.DB
}

func NewParkingLotRepositoryImpl(Db *gorm.DB) ParkingLotRepository {
	return &ParkingLotRepositoryImpl{Db: Db}
}

func (d *ParkingLotRepositoryImpl) Save(tx *gorm.DB, parkingLot model.ParkingLots) (int64, error) {
	if err := tx.Create(&parkingLot).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return parkingLot.LotID, nil
}

func (d *ParkingLotRepositoryImpl) GetParkingLotStatus() (resp []response.ParkingLotStatusResponse, err error) {
	// Query the database to fetch parking lot status
	rows, err := d.Db.Raw(`
					SELECT pl.lot_id, pl.name AS lot_name, ps.slot_id, ps.slot_no, ps.is_available, ps.in_maintenance, pf.vehicle_number 
					FROM parking_lots pl
					JOIN parking_slots ps ON pl.lot_id = ps.lot_id
					LEFT JOIN parking_fees pf ON ps.slot_id = pf.slot_id
					WHERE pf.parking_end_time = ?  					      
				`, data.DEFAULT_PARK_END_TIME).Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Iterate over the rows and populate the parkingLotStatus slice
	for rows.Next() {
		var response response.ParkingLotStatusResponse
		var vehicleNumber sql.NullString
		err := rows.Scan(&response.LotID, &response.LotName, &response.SlotID, &response.SlotNumber, &response.IsAvailable, &response.InMaintenance, &vehicleNumber)
		if err != nil {
			return nil, err
		}
		if vehicleNumber.Valid {
			response.VehicleNumber = vehicleNumber.String
		} else {
			response.VehicleNumber = "" // Assign an empty string if vehicle_number is NULL
		}
		resp = append(resp, response)
	}

	// Check for any errors occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return resp, nil
}
