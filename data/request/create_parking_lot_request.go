package request

type CreateParkingLotRequest struct {
	Name       string `json:"name"`
	TotalSlots int    `json:"total_slots"`
}
