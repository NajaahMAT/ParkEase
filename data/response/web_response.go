package response

import "time"

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type DailySummaryResponse struct {
	TotalVehiclesParked int    `json:"totalVehiclesParked"`
	TotalParkingTime    string `json:"totalParkingTime"`
	TotalFeeCollected   int    `json:"totalFeeCollected"`
}

// Response struct
type AvailableSlotsResponse struct {
	Lots []LotSlots `json:"lots"`
}

type LotSlots struct {
	LotID int64   `json:"lot_id"`
	Slots []int64 `json:"slots"`
}

type UnParkVehicleResponse struct {
	BillID       int64     `json:"bill_id"`
	VehicleNo    string    `json:"vehicle_no"`
	ParkingStart time.Time `json:"parking_start"`
	ParkingEnd   time.Time `json:"parking_end"`
	ParkingFee   float64   `json:"parking_fee"`
}

type ParkingLotStatusResponse struct {
	LotID         int64  `json:"lot_id"`
	LotName       string `json:"lot_name"`
	SlotID        int64  `json:"slot_id"`
	SlotNumber    int    `json:"slot_number"`
	IsAvailable   bool   `json:"is_available"`
	InMaintenance bool   `json:"in_maintenance"`
	VehicleNumber string `json:"vehicle_number,omitempty"`
}

// ParkingStatsResponse represents the response structure for parking stats
type ParkingStatsResponse struct {
	TotalVehicles    int64         `json:"totalVehicles"`
	TotalParkingTime time.Duration `json:"totalParkingTime"`
	TotalParkingFee  float64       `json:"totalParkingFee"`
}
