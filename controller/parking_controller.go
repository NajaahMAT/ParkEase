package controller

import (
	"ParkEase/data/request"
	"ParkEase/data/response"
	"ParkEase/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ParkingController struct {
	parkingService service.ParkingService
}

func NewParkingController(service service.ParkingService) *ParkingController {
	return &ParkingController{
		parkingService: service,
	}
}

func (controller *ParkingController) Create(ctx *gin.Context) {
	createLotRequest := request.CreateParkingLotRequest{}
	err := ctx.ShouldBindJSON(&createLotRequest)
	if err != nil {
		// Handle JSON binding error
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	lotID, err := controller.parkingService.Create(createLotRequest)
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
			"lot_id": lotID,
		},
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ParkingController) GetAvailableSlotsByLot(ctx *gin.Context) {
	slotsResponse, err := controller.parkingService.GetAvailableSlotsByLot()
	if err != nil {
		// Handle error from service layer
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	// Create a response struct
	response := response.AvailableSlotsResponse{
		Lots: slotsResponse.Lots,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}

func (controller *ParkingController) ParkVehicle(ctx *gin.Context) {
	parkRequest := request.ParkVehicleRequest{}
	err := ctx.ShouldBindJSON(&parkRequest)
	if err != nil {
		// Handle JSON binding error
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	feeID, err := controller.parkingService.ParkVehicle(parkRequest)
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
			"parking_fee_id": feeID,
		},
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ParkingController) UnParkVehicle(ctx *gin.Context) {
	parkingFeeID := ctx.Param("parkingFeeID")
	id, err := strconv.ParseInt(parkingFeeID, 10, 64)
	if err != nil {
		// Handle error from service layer
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Error",
			Data:   err.Error(),
		})
		return
	}

	parkFeeInfo, err := controller.parkingService.UnParkVehicle(id)
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
		Data: map[string]interface{}{
			"bill_info": parkFeeInfo,
		},
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ParkingController) GetParkingLotStatus(ctx *gin.Context) {
	parkingLotStatus, err := controller.parkingService.GetParkingLotStatus()
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
		Data: map[string]interface{}{
			"LotStatus": parkingLotStatus,
		},
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ParkingController) GetParkingStats(ctx *gin.Context) {
	dateString := ctx.Query("date")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	parkingStats, err := controller.parkingService.GetParkingStats(date)
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
		Data:   parkingStats,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
