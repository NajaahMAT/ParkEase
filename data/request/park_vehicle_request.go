package request

type ParkVehicleRequest struct {
	SlotID          int64  `json:"slot_id"`
	VehicleNumber   string `json:"vehicle_number" validate:"required"`
	IsSlotAvailable bool   `json:"is_slot_available"`
}
