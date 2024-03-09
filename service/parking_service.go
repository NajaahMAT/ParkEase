package service

import (
	"ParkEase/data/request"
	"ParkEase/data/response"
	"ParkEase/model"
	"ParkEase/repository"
	"math"
	"time"

	validator "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ParkingServiceImpl struct {
	ParkLotRepository  repository.ParkingLotRepository
	ParkSlotRepository repository.ParkingSlotRepository
	ParkFeeRepository  repository.ParkingFeeRepository
	Validate           *validator.Validate
	Db                 *gorm.DB // Add Db field
}

func NewParkingServiceImpl(db *gorm.DB, parkLotRepository repository.ParkingLotRepository,
	parkSlotRepository repository.ParkingSlotRepository, parkFeeRepository repository.ParkingFeeRepository, validate *validator.Validate) ParkingService {
	return &ParkingServiceImpl{
		ParkLotRepository:  parkLotRepository,
		ParkSlotRepository: parkSlotRepository,
		ParkFeeRepository:  parkFeeRepository,
		Validate:           validate,
		Db:                 db, // Assign the Db instance
	}
}

func (t *ParkingServiceImpl) Create(req request.CreateParkingLotRequest) (lotID int64, err error) {
	err = t.Validate.Struct(req)
	if err != nil {
		return
	}

	// Start a transaction
	tx := t.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create new parking lot entry
	lotModel := model.ParkingLots{
		Name:       req.Name,
		TotalSlots: req.TotalSlots,
	}

	lotID, err = t.ParkLotRepository.Save(tx, lotModel)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Create parking slots associated with the parking lot
	for i := 1; i <= req.TotalSlots; i++ {
		slot := model.ParkingSlots{
			LotID:         lotID,
			SlotNo:        i,
			InMaintenance: false, // Assuming slots are not in maintenance initially
		}
		_, err := t.ParkSlotRepository.Save(tx, slot)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return lotID, nil
}

func (t *ParkingServiceImpl) GetAvailableSlotsByLot() (resp response.AvailableSlotsResponse, err error) {
	availableSlots, err := t.ParkSlotRepository.GetAvailableSlotsByLot()
	if err != nil {
		return resp, err
	}

	// Convert available slots to the desired response struct
	var slotsByLot = make(map[int64][]int64)
	for _, slot := range availableSlots {
		slotsByLot[slot.LotID] = append(slotsByLot[slot.LotID], int64(slot.SlotNo))
	}

	var lots []response.LotSlots
	for lotID, slots := range slotsByLot {
		lots = append(lots, response.LotSlots{
			LotID: lotID,
			Slots: slots,
		})
	}

	return response.AvailableSlotsResponse{Lots: lots}, nil
}

func (t *ParkingServiceImpl) ParkVehicle(req request.ParkVehicleRequest) (feeID int64, err error) {
	err = t.Validate.Struct(req)
	if err != nil {
		return 0, err
	}

	// Start a transaction
	tx := t.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create new parking vehicle entry in the parking fee table
	parkFeeModel := model.ParkingFees{
		SlotID:           req.SlotID,
		VehicleNumber:    req.VehicleNumber,
		ParkingStartTime: time.Now(),
		//ParkingEndTime:   time.Now().Add(time.Hour), //By Default Parking End Time is calculated as 1 hour from start time, but it will be recalculated while unparking.
	}
	feeID, err = t.ParkFeeRepository.Save(tx, parkFeeModel)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	req.IsSlotAvailable = false //to make sure the park slot availability status false

	// update the slot available status to false in parking slots table
	err = t.ParkSlotRepository.UpdateSlotAvailableStatus(tx, req)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return feeID, nil
}

func (t *ParkingServiceImpl) UnParkVehicle(id int64) (resp response.UnParkVehicleResponse, err error) {

	// Retrieve the parking fee record based on the provided ID
	parkingData, err := t.ParkFeeRepository.GetParkingFeeByID(id)
	if err != nil {
		return resp, err
	}

	// Calculate the duration for which the vehicle was parked
	current := time.Now()
	parkingFee := calculateParkingFee(parkingData.ParkingStartTime, current)

	// Create new parking vehicle entry in the parking fee table
	parkingData.ParkingEndTime = current
	parkingData.ParkingFee = float64(parkingFee)

	// Start a transaction
	tx := t.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err = t.ParkFeeRepository.UpdateParkingFees(tx, parkingData)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	parkSlotReq := request.ParkVehicleRequest{
		SlotID:          parkingData.SlotID,
		VehicleNumber:   parkingData.VehicleNumber,
		IsSlotAvailable: true,
	}
	// update the slot available status to true in parking slots table
	err = t.ParkSlotRepository.UpdateSlotAvailableStatus(tx, parkSlotReq)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return resp, err
	}

	resp.BillID = parkingData.ID
	resp.VehicleNo = parkingData.VehicleNumber
	resp.ParkingStart = parkingData.ParkingStartTime
	resp.ParkingEnd = parkingData.ParkingEndTime
	resp.ParkingFee = parkingData.ParkingFee

	return resp, nil
}

func calculateParkingFee(parkedAt time.Time, parkedEnd time.Time) int {
	duration := parkedEnd.Sub(parkedAt)
	hours := int(math.Ceil(duration.Hours())) //even if the duration is less than an hour, it will be rounded up to one hour. For example, if a vehicle is parked for 1 hour and 5 minutes, it will be considered as 2 hours
	if hours == 0 {
		hours = 1 // Minimum parking fee is for one hour
	}

	return hours * 10

}

func (t *ParkingServiceImpl) GetParkingLotStatus() (resp []response.ParkingLotStatusResponse, err error) {

	parkinglotStatus, err := t.ParkLotRepository.GetParkingLotStatus()
	if err != nil {
		return resp, err
	}

	return parkinglotStatus, nil
}

func (t *ParkingServiceImpl) GetParkingStats(date time.Time) (resp response.ParkingStatsResponse, err error) {

	parkingStats, err := t.ParkFeeRepository.GetParkingStats(date)
	if err != nil {
		return resp, err
	}

	return parkingStats, nil
}
