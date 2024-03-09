package request

type SlotMaintenanceRequest struct {
	SlotID        int64  `json:"slot_id"`
	Reason        string `json:"reason"`
	InMaintenance bool   `json:"in_maintenance"`
}

// type SlotRestoreFromMaintenanceRequest struct {
// 	SlotID        int64 `json:"slot_id"`
// 	MaintenanceID int64 `json:"maintenance_id"`
// 	InMaintenance bool  `json:"in_maintenance"`
// }
