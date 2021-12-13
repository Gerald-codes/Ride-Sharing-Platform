package TripDB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Declare Variables and Structure
var id string

type Trip struct { // map this type to the record in the table
	TripID            int
	PassengerID       int
	DriverID          int
	PickUpPostalCode  int
	DropOffPostalCode int
	TripStartTime     string
	TripEndTime       string
	Status            string
}

type TripDetails struct { // map this type to the record in the table
	TripID            int
	FirstName         string
	LastName          string
	MobileNo          int
	PickUpPostalCode  int
	DropOffPostalCode int
}

// Get All Trip Records
func GetRecords() (res []Trip) {
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, Qerr := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ride_sharing")

	// handle error
	if Qerr != nil {
		panic(Qerr.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()
	results, err := db.Query("Select * FROM ride_sharing.Trips")
	if err != nil {
		panic(err.Error())
	}
	// Declare Trip array to store multiple columns of trips
	var TripRecords []Trip
	for results.Next() {
		// map this type to the record in the table
		var tripDetails Trip
		err = results.Scan(&tripDetails.TripID, &tripDetails.PassengerID, &tripDetails.DriverID, &tripDetails.Status, &tripDetails.PickUpPostalCode,
			&tripDetails.DropOffPostalCode, &tripDetails.TripStartTime, &tripDetails.TripEndTime)
		if err != nil {
			panic(err.Error())
		}
		// Append trip details into the array of Trip records
		TripRecords = append(TripRecords, tripDetails)
	}
	// fmt.Print("ARRAY!!", TripRecords)
	return TripRecords
}

// GET latest ID
func GetLatestID() (res string) {
	db, Qerr := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ride_sharing")

	// handle error
	if Qerr != nil {
		panic(Qerr.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()

	results, err := db.Query("SELECT COALESCE(MAX(TripID),0) FROM ride_sharing.Trips")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var res string
		err = results.Scan(&res)
		if err != nil {
			panic(err.Error())
		}
		id = res
	}
	return id
}

// Passenger Create Trip Record
func InsertRecord(db *sql.DB, TID int, PID int, PU int, DO int, S string) {
	query := fmt.Sprintf("INSERT INTO ride_sharing.Trips (TripID, PassengerID, PickUpPostalCode, DropOffPostalCode, Status) VALUES ( %d, %d, %d, %d,'%s')",
		TID, PID, PU, DO, S)
	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

// Driver Get all pending records for selection
func GetPendingRecords(DID int) (res []TripDetails) {
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, Qerr := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ride_sharing")

	// handle error
	if Qerr != nil {
		panic(Qerr.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()
	// Get data from both Trip and passenger user inner join
	query := fmt.Sprintf(
		"SELECT t.TripID, p.FirstName,p.LastName,p.MobileNo,t.PickUpPostalCode, t.DropOffPostalCode FROM ride_sharing.Passengers p INNER JOIN ride_sharing.Trips t ON p.PassengerID=t.PassengerID WHERE t.DriverID=%d AND t.Status='Pending' ",
		DID)

	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	// Declare TripDetails array to store multiple columns of trips
	var pendingTripRecords []TripDetails
	for results.Next() {
		// map this type to the record in the table
		var tripDetails TripDetails
		err = results.Scan(&tripDetails.TripID, &tripDetails.FirstName, &tripDetails.LastName, &tripDetails.MobileNo, &tripDetails.PickUpPostalCode,
			&tripDetails.DropOffPostalCode)
		if err != nil {
			panic(err.Error())
		}
		// Append trip details into the array of Trip records
		pendingTripRecords = append(pendingTripRecords, tripDetails)
	}
	// fmt.Print("ARRAY!!", pendingTripRecords)
	return pendingTripRecords
}

// Driver Initiate Start Trip
func StartTrip(db *sql.DB, ST string, S string, TID int) {
	query := fmt.Sprintf(
		"UPDATE ride_sharing.Trips SET TripStartTime='%s', Status='%s' WHERE TripID=%d ",
		ST, S, TID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

// Driver Initiate End Trip
func EndTrip(db *sql.DB, ET string, S string, TID int) {
	query := fmt.Sprintf(
		"UPDATE ride_sharing.Trips SET TripEndTime='%s', Status='%s' WHERE TripID=%d",
		ET, S, TID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

// Passenger Get Trip History
func GetTripHistory(PID int) (res []Trip) {
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, Qerr := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ride_sharing")

	// handle error
	if Qerr != nil {
		panic(Qerr.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()

	query := fmt.Sprintf(
		"Select * FROM ride_sharing.Trips WHERE Status='Completed' AND PassengerID=%d ORDER BY TripID DESC", PID)

	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	// Declare Trip array to store multiple columns of trips
	var TripRecords []Trip
	for results.Next() {
		// map this type to the record in the table
		var tripDetails Trip
		err = results.Scan(&tripDetails.TripID, &tripDetails.PassengerID, &tripDetails.DriverID, &tripDetails.Status, &tripDetails.PickUpPostalCode,
			&tripDetails.DropOffPostalCode, &tripDetails.TripStartTime, &tripDetails.TripEndTime)
		if err != nil {
			panic(err.Error())
		}

		// Append trip details into the array of Trip records
		TripRecords = append(TripRecords, tripDetails)
	}
	// fmt.Print("ARRAY!!", TripRecords)
	return TripRecords
}

// Main
func TripDB(method string, t Trip) {
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ride_sharing")

	// handle error
	if err != nil {
		panic(err.Error())
	}
	if method == "Insert" {
		InsertRecord(db, t.TripID, t.PassengerID, t.PickUpPostalCode, t.DropOffPostalCode, t.Status)
		fmt.Println("Insert Trip", t.TripStartTime, " Database")
	} else if method == "Start" {
		fmt.Println("Started Trip", t.TripID, " Database")
		StartTrip(db, t.TripStartTime, t.Status, t.TripID)
	} else if method == "End" {
		EndTrip(db, t.TripEndTime, t.Status, t.TripID)
		fmt.Println("Ended Trip", t.TripID, " Database")
	}

	// defer the close till after the main function has finished executing
	defer db.Close()

	fmt.Println("Database opened")
}
