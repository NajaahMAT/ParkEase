# ParkEase
ParkEase is a parking lot management system developed in Go, designed to efficiently manage parking spaces in a parking lot. It provides a set of APIs for parking managers and users to interact with the system, enabling functionalities such as creating parking lots, parking vehicles, unparking vehicles, viewing parking lot status, putting slots in maintenance mode, and retrieving daily summaries of parking activities.

## Features of the Application
* Create Parking Lot: Parking managers can create parking lots with desired parking spaces/slots.
* Park Vehicle: Users (Vehicle owners) can park their vehicles in the nearest available parking slot in a chosen parking lot.
* Unpark Vehicle: Users can unpark their vehicles, calculating parking fees based on the parking duration.
* View Parking Lot Status: Parking managers can view the current status of parking lots, including which cars are parked in which slots.
* Maintenance Mode: Parking managers can put any parking space/slot into maintenance mode and back to working state at any time.
* Daily Summary: Parking managers can retrieve total number of vehicles parked on any day, total parking time, and the total fee collected on that day.

### To Run the Project:
1. Clone the go project. using the following command:   `git clone https://github.com/NajaahMAT/ParkEase.git`
2. Update go modules.  `go mod tidy`
4. Establish the Database Connection , Provide the following Params 
   
    ```
    const (
        user     = "root"       // change as per your MySQL user
        password = "password" // change as per your MySQL password
        dbName   = "park_ease"
        host     = "localhost"
        port     = 3306
    )
    ```
Database Configurations Should be edited in the "TripManagementSystem/config/mysql_database.go" file
5. To Build the project.  `go build`
6. To execute the project. `go run .`

### Assumptions:
a.While Calculating parking fee even if the duration is less than an hour, it will be rounded up to one hour. For example, if a vehicle is parked for 1 hour and 5 minutes, it will be considered as 2 hours
b. In here we are not considering the lots seperation according to Vehicle Type, (Eg, Car/ Bike/ Bus etc)
c. While creating slots, assuming that slots are not in maintenance initially.

### Database Design:
The Tables , Relationships and Key constraints are created through AutoMigration.For the file path "ParkEase/helper/db_migrations.go"
1. To Store Park lots Data: `parking_lots`
2. To Store Park Slots Data: `parking_slots`
3. To Store Parking Fee Calculation Related Data: `parking_fees`
4. To store Parking Slot Maintenance Data: `parking_slot_maintenances`

| Table Name             | Column Name       | Data Type     | Constraints                                     | Description                                      |
|------------------------|-------------------|---------------|-------------------------------------------------|--------------------------------------------------|
| parking_fees           | ID                | int64         | primary_key, auto_increment                     | Unique identifier for the ParkingFee            |
|                        | SlotID            | int64         | not null                                        | Parking slot identifier                         |
|                        | VehicleNumber     | varchar(50)   | type:varchar(50), not null                      | Vehicle number parked                           |
|                        | ParkingStartTime  | time.Time     | not null                                        | Start time of parking                           |
|                        | ParkingEndTime    | time.Time     | default:9999-12-31 23:59:59                    | End time of parking (default: 31 Dec 9999)      |
|                        | ParkingFee        | decimal(10,6) | type:decimal(10,6)                              | Parking fee                                     |
| parking_lots           | LotID             | int64         | primary_key, auto_increment, comment:Unique identifier for the ParkingLot | Unique identifier for the ParkingLot |
|                        | Name              | varchar(50)   | type:varchar(50), not null, comment:Name of the Parking Lot | Name of the Parking Lot                    |
|                        | TotalSlots        | int           | type:int, not null, comment:Total slots in the parking lot | Total slots in the parking lot              |
|                        | CreatedDate       | time.Time     | autoCreateTime                                 | Automatically generated creation timestamp   |
|                        | UpdatedDate       | time.Time     | autoUpdateTime                                 | Automatically generated update timestamp     |
| parking_slot_maintenances | MaintenanceID   | int64         | primary_key, auto_increment, comment:Unique identifier for the Parking Slot Maintenance | Unique identifier for the Parking Slot Maintenance |
|                            | SlotID         | int64         | not null, comment:Parking LotID               | Parking slot identifier                       |
|                            | MaintenanceStart | time.Time   | not null                                      | Start time of maintenance                     |
|                            | MaintenanceEnd   | time.Time   | default:9999-12-31 23:59:59                  | End time of maintenance (default: 31 Dec 9999) |
|                            | Reason          | varchar(200) | type:varchar(200), not null                   | Reason for maintenance                         |
| parking_slots          | SlotID            | int64         | primary_key, auto_increment, comment:Unique identifier for the Parking Slot | Unique identifier for the Parking Slot   |
|                        | LotID             | int64         | type:int, not null, comment:Parking LotID      | Parking lot identifier                        |
|                        | SlotNo            | int           | type:int, not null                             | Parking slot number                           |
|                        | InMaintenance     | boolean       | type:boolean, default:false                    | Maintenance status of the slot                |
|                        | IsAvailable       | boolean       | type:boolean, default:true                     | Availability status of the slot               |
|                        | CreatedDate       | time.Time     | autoCreateTime                                 | Automatically generated creation timestamp   |
|                        | UpdatedDate       | time.Time     | autoUpdateTime                                 | Automatically generated update timestamp     |


### Api Details.

| API Name                  | Description                                             | HTTP Method | Request Struct                 | Response Struct               |
|---------------------------|---------------------------------------------------------|-------------|--------------------------------|-------------------------------|
| Create Parking Lot        | Create a new parking lot with a given name and slots   | POST        | { "name": string, "total_slots": int } | { "code": int, "status": string, "data": { "lot_id": int } } |
| Get Available Parking Slots By Lots | Retrieve available parking slots grouped by lots    | GET         | None                           | { "lots": [ { "lot_id": int, "slots": [int] } ] } |
| Park Vehicle              | Park a vehicle in a specific parking slot              | POST        | { "slot_id": int, "vehicle_number": string } | { "code": int, "status": string, "data": { "parking_fee_id": int } } |
| Unpark Vehicle            | Unpark a vehicle from a specific parking slot         | PUT         | None                           | { "code": int, "status": string, "data": { "bill_info": { "bill_id": int, "vehicle_no": string, "parking_start": string, "parking_end": string, "parking_fee": float64 } } } |
| Get Parking Lots Status  | Retrieve the status of parking slots in a parking lot | GET         | None                           | { "code": int, "status": string, "data": { "LotStatus": [ { "lot_id": int, "lot_name": string, "slot_id": int, "slot_number": int, "is_available": bool, "in_maintenance": bool } ] } } |
| Put Slots In Maintenance | Put parking slots into maintenance mode                | POST        | { "slot_id": int, "reason": string, "in_maintenance": bool } | { "code": int, "status": string, "data": { "maintenance_id": int } } |
| Restore Slots From Maintenance | Restore parking slots from maintenance mode       | PUT         | { "slot_id": int, "reason": string, "in_maintenance": bool } | { "code": int, "status": string } |
| Get Parking Statistics   | Retrieve parking statistics for a given date          | GET         | None                           | { "code": int, "status": string, "data": { "totalVehicles": int, "totalParkingTime": int, "totalParkingFee": float64 } } |

