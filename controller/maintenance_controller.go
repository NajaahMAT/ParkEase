package controller

import (
	"ParkEase/data/request"
	"ParkEase/data/response"
	"ParkEase/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MaintenanceController struct {
	maintenanceService service.MaintenanceService
}

func NewMaintenanceController(service service.MaintenanceService) *MaintenanceController {
	return &MaintenanceController{
		maintenanceService: service,
	}
}

func (controller *MaintenanceController) PutSlotIntoMaintenance(ctx *gin.Context) {
	maintenanceRequest := request.SlotMaintenanceRequest{}
	err := ctx.ShouldBindJSON(&maintenanceRequest)
	if err != nil {
		// Handle JSON binding error
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	maintenaceID, err := controller.maintenanceService.PutSlotsIntoMaintenance(maintenanceRequest)
	if err != nil {
		// Handle error from service layer
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data: map[string]int64{
			"maintenance_id": maintenaceID,
		},
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MaintenanceController) RestoreSlotFromMaintenance(ctx *gin.Context) {
	maintenanceRequest := request.SlotMaintenanceRequest{}
	err := ctx.ShouldBindJSON(&maintenanceRequest)
	if err != nil {
		// Handle JSON binding error
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	maintenaceID := ctx.Param("maintenaceID")
	id, err := strconv.ParseInt(maintenaceID, 10, 64)
	if err != nil {
		// Handle error from service layer
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	err = controller.maintenanceService.RestoreSlotsFromMaintenance(maintenanceRequest, id)
	if err != nil {
		// Handle error from service layer
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
