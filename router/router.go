package router

import (
	"ParkEase/config"
	"ParkEase/controller"
	"ParkEase/helper"
	"ParkEase/repository"
	"ParkEase/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Dependencies struct {
	ParkingController     *controller.ParkingController
	MaintenanceController *controller.MaintenanceController
}

func InitializeDependencies() *Dependencies {
	db := config.DatabaseConnection()
	validate := validator.New()

	// Call AutoMigrate from migrations.go
	if err := helper.AutoMigrate(db); err != nil {
		// Handle error, for example, log it or panic
		log.Fatalf("Migration failed: %v", err)
	}

	// Setup for Parking
	parkingLotRepository := repository.NewParkingLotRepositoryImpl(db)
	parkingSlotRepository := repository.NewParkingSlotRepositoryImpl(db)
	parkingFeeRepository := repository.NewParkingFeeRepositoryImpl(db)
	parkingService := service.NewParkingServiceImpl(parkingLotRepository, parkingSlotRepository, parkingFeeRepository, validate)
	parkingController := controller.NewParkingController(parkingService)

	//Setup for Maintenance
	parkingSlotMaintenanceRepository := repository.NewParkingSlotMaintenanceRepositoryImpl(db)
	maintenanceService := service.NewMaintenanceServiceImpl(parkingSlotRepository, parkingSlotMaintenanceRepository, validate)
	maintenanceController := controller.NewMaintenanceController(maintenanceService)

	return &Dependencies{
		ParkingController:     parkingController,
		MaintenanceController: maintenanceController,
	}
}

func NewRouter(deps *Dependencies) *gin.Engine {
	router := gin.Default()

	baseRouter := router.Group("/api")
	baseRouter.POST("/park", deps.ParkingController.ParkVehicle)
	baseRouter.PUT("/unpark/:parkingFeeID", deps.ParkingController.UnParkVehicle)
	baseRouter.GET("/parking/stats", deps.ParkingController.GetParkingStats)

	parkingLotsRouter := baseRouter.Group("/parking/lots")
	parkingLotsRouter.POST("", deps.ParkingController.Create)
	parkingLotsRouter.GET("/slots", deps.ParkingController.GetAvailableSlotsByLot)
	parkingLotsRouter.GET("/status", deps.ParkingController.GetParkingLotStatus)

	parkingSlotsRouter := baseRouter.Group("/parking/slot")
	parkingSlotsRouter.POST("/maintenance", deps.MaintenanceController.PutSlotIntoMaintenance)
	parkingSlotsRouter.PUT("/restore/:maintenaceID", deps.MaintenanceController.RestoreSlotFromMaintenance)

	return router
}
