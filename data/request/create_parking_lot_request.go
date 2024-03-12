package request

type CreateParkingLotRequest struct {
	Name       string `json:"name" validate:"required"`
	TotalSlots int    `json:"total_slots" validate:"required,min=1"`
}
