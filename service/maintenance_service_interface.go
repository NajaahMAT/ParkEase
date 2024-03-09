package service

import "ParkEase/data/request"

type MaintenanceService interface {
	PutSlotsIntoMaintenance(req request.SlotMaintenanceRequest) (maintenanceID int64, err error)
	RestoreSlotsFromMaintenance(req request.SlotMaintenanceRequest, maintenanceID int64) (err error)
}
