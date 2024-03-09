package repository

import (
	"ParkEase/data/response"
	"ParkEase/model"
	"errors"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ParkingFeeRepositoryImpl struct {
	Db *gorm.DB
}

func NewParkingFeeRepositoryImpl(Db *gorm.DB) ParkingFeeRepository {
	return &ParkingFeeRepositoryImpl{Db: Db}
}

func (d *ParkingFeeRepositoryImpl) Save(tx *gorm.DB, parkingFee model.ParkingFees) (int64, error) {
	if err := tx.Create(&parkingFee).Error; err != nil {
		return 0, err
	}

	return int64(parkingFee.ID), nil
}

func (t *ParkingFeeRepositoryImpl) GetParkingFeeByID(id int64) (resp model.ParkingFees, err error) {

	// Retrieve the parking fee record based on the provided ID
	parkingFee := model.ParkingFees{}
	if err := t.Db.First(&parkingFee, id).Error; err != nil {
		return resp, err
	}

	return parkingFee, nil
}

func (t *ParkingFeeRepositoryImpl) UpdateParkingFees(tx *gorm.DB, req model.ParkingFees) error {
	// Prepare a map of columns to update
	updates := map[string]interface{}{
		"parking_end_time": req.ParkingEndTime,
		"parking_fee":      req.ParkingFee,
	}

	// Proceed with updating the status
	if err := tx.Model(&model.ParkingFees{}).Where("id = ?", req.ID).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

func (d *ParkingFeeRepositoryImpl) GetParkingStats(date time.Time) (resp response.ParkingStatsResponse, err error) {
	var parkingStats response.ParkingStatsResponse

	// Calculate total number of vehicles parked on the given date
	if err := d.Db.Model(&model.ParkingFees{}).Where("DATE(parking_start_time) = ?", date.Format("2006-01-02")).Count(&parkingStats.TotalVehicles).Error; err != nil {
		return resp, err
	}

	// Calculate total parking time and total fee collected on the given date
	var totalDuration time.Duration
	var totalFee float64
	rows, err := d.Db.Table("parking_fees").
		Select("TIMEDIFF(parking_end_time, parking_start_time) AS duration, parking_fee").
		Where("DATE(parking_start_time) = ?", date.Format("2006-01-02")).Rows()
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var durationStr string
		var fee float64
		rows.Scan(&durationStr, &fee)

		parsedDuration, err := parseDuration(durationStr)
		if err != nil {
			return resp, err
		}
		totalDuration += parsedDuration
		totalFee += fee
	}

	parkingStats.TotalParkingTime = totalDuration / time.Second // Convert to seconds
	parkingStats.TotalParkingFee = totalFee

	return parkingStats, nil
}

// parseDuration parses a duration string in the format "hh:mm:ss.ms" into a time.Duration object.
func parseDuration(durationStr string) (time.Duration, error) {
	parts := strings.Split(durationStr, ":")
	if len(parts) != 3 {
		return 0, errors.New("invalid duration format")
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}
	secondsWithMS := strings.Split(parts[2], ".")
	seconds, err := strconv.Atoi(secondsWithMS[0])
	if err != nil {
		return 0, err
	}
	milliseconds, err := strconv.Atoi(secondsWithMS[1])
	if err != nil {
		return 0, err
	}

	return time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second + time.Duration(milliseconds)*time.Millisecond, nil
}
