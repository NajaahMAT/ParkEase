package request

type ParkVehicleRequest struct {
	SlotID          int64  `json:"slot_id" validate:"required,min=1"`
	VehicleNumber   string `json:"vehicle_number" validate:"required,alphanum"`
	IsSlotAvailable bool   `json:"is_slot_available"`
}
