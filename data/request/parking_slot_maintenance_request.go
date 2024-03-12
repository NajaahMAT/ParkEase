package request

type SlotMaintenanceRequest struct {
	SlotID        int64  `json:"slot_id" validate:"required,min=1"`
	Reason        string `json:"reason"`
	InMaintenance bool   `json:"in_maintenance"`
}
