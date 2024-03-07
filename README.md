# ParkEase
ParkEase is a parking lot management system developed in Go, designed to efficiently manage parking spaces in a parking lot. It provides a set of APIs for parking managers and users to interact with the system, enabling functionalities such as creating parking lots, parking vehicles, unparking vehicles, viewing parking lot status, putting slots in maintenance mode, and retrieving daily summaries of parking activities.

## Features of the Application
* Create Parking Lot: Parking managers can create parking lots with desired parking spaces/slots.
* Park Vehicle: Users (Vehicle owners) can park their vehicles in the nearest available parking slot in a chosen parking lot.
* Unpark Vehicle: Users can unpark their vehicles, calculating parking fees based on the parking duration.
* View Parking Lot Status: Parking managers can view the current status of parking lots, including which cars are parked in which slots.
* Maintenance Mode: Parking managers can put any parking space/slot into maintenance mode and back to working state at any time.
* Daily Summary: Parking managers can retrieve total number of vehicles parked on any day, total parking time, and the total fee collected on that day.