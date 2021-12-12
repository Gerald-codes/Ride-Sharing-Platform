package TripDB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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

func StartTrip(db *sql.DB, DID int, ST string, S string, TID int) {
	query := fmt.Sprintf(
		"UPDATE ride_sharing.Trips SET DriverID=%d, TripStartTime='%s', Status='%s' WHERE TripID=%d ",
		DID, ST, S, TID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func EndTrip(db *sql.DB, ET string, S string, TID int) {
	query := fmt.Sprintf(
		"UPDATE ride_sharing.Trips SET TripEndTime='%s', Status='%s' WHERE TripID=%d",
		ET, S, TID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func InsertRecord(db *sql.DB, TID int, PID int, PU int, DO int, S string) {
	query := fmt.Sprintf("INSERT INTO ride_sharing.Trips (TripID, PassengerID, PickUpPostalCode, DropOffPostalCode, Status) VALUES ( %d, %d, %d, %d,'%s')",
		TID, PID, PU, DO, S)
	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func GetPendingRecords() (res []TripDetails) {
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, Qerr := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ride_sharing")

	// handle error
	if Qerr != nil {
		panic(Qerr.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()

	results, err := db.Query("SELECT t.TripID, p.FirstName,p.LastName,p.MobileNo,t.PickUpPostalCode, t.DropOffPostalCode FROM ride_sharing.Passengers p INNER JOIN ride_sharing.Trips t ON p.PassengerID=t.PassengerID")

	if err != nil {
		panic(err.Error())
	}
	var pendingTripRecords []TripDetails
	for results.Next() {
		// map this type to the record in the table
		var tripDetails TripDetails
		err = results.Scan(&tripDetails.TripID, &tripDetails.FirstName, &tripDetails.LastName, &tripDetails.MobileNo, &tripDetails.PickUpPostalCode,
			&tripDetails.DropOffPostalCode)
		if err != nil {
			panic(err.Error())
		}
		// fmt.Println(tripDetails.TripID, tripDetails.FirstName, tripDetails.LastName, tripDetails.MobileNo, tripDetails.PickUpPostalCode,
		// 	tripDetails.DropOffPostalCode)
		pendingTripRecords = append(pendingTripRecords, tripDetails)
	}
	// fmt.Print("ARRAY!!", pendingTripRecords)
	return pendingTripRecords
}

func GetRecords() (res []Trip) {
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, Qerr := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ride_sharing")

	// handle error
	if Qerr != nil {
		panic(Qerr.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()
	results, err := db.Query("Select * FROM ride_sharing.Trips WHERE Status='Completed'")
	if err != nil {
		panic(err.Error())
	}
	var TripRecords []Trip
	for results.Next() {
		// map this type to the record in the table
		var tripDetails Trip
		err = results.Scan(&tripDetails.TripID, &tripDetails.PassengerID, &tripDetails.DriverID, &tripDetails.Status, &tripDetails.PickUpPostalCode,
			&tripDetails.DropOffPostalCode, &tripDetails.TripStartTime, &tripDetails.TripEndTime)
		if err != nil {
			panic(err.Error())
		}
		TripRecords = append(TripRecords, tripDetails)
	}
	// fmt.Print("ARRAY!!", TripRecords)
	return TripRecords
}

func GetLatestID() (res string) {
	db, Qerr := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ride_sharing")

	// handle error
	if Qerr != nil {
		panic(Qerr.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()

	results, err := db.Query("SELECT MAX(TripID) FROM ride_sharing.Trips")

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
		StartTrip(db, 18, t.TripStartTime, t.Status, 9)
	} else if method == "End" {
		EndTrip(db, t.TripStartTime, t.Status, t.TripID)
		fmt.Println("Ended Trip", t.TripID, " Database")
	}

	// defer the close till after the main function has finished executing
	defer db.Close()

	fmt.Println("Database opened")
}
