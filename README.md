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