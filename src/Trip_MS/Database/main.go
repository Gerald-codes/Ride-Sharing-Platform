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

func EditRecord(db *sql.DB, ID int, FN string, LN string, EA string) {
	query := fmt.Sprintf(
		"UPDATE Passenger SET FirstName='%s', LastName='%s', EmailAddress='%s' WHERE PassengerID=%d",
		FN, LN, EA, ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func InsertRecord(db *sql.DB, TID int, PID int, PU int, DO int, ST string, S string) {
	query := fmt.Sprintf("INSERT INTO ride_sharing.Trips (TripID, PassengerID, PickUpPostalCode, DropOffPostalCode, TripStartTime, Status) VALUES ( %d, %d, %d, %d,'%s','%s')",
		TID, PID, PU, DO, ST, S)
	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func GetRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM ride_sharing.Passenger")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var trip Trip
		err = results.Scan(&trip.TripID, &trip.PassengerID, &trip.DriverID, &trip.PickUpPostalCode,
			&trip.DropOffPostalCode, &trip.TripStartTime, &trip.TripEndTime, &trip.Status)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(trip.TripID, trip.PassengerID, trip.DriverID, trip.PickUpPostalCode,
			trip.DropOffPostalCode, trip.TripStartTime, trip.TripEndTime, trip.Status)
	}
}
func GetLatestID() (res string) {
	// Use mysql as driverName and a valid DSN as dataSourceName:
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
		InsertRecord(db, t.TripID, t.PassengerID, t.PickUpPostalCode, t.DropOffPostalCode, t.TripStartTime, t.Status)
		fmt.Println("Insert Trip", t.TripStartTime, " Database")
	} else if method == "Update" {
		fmt.Println("Updated Trip", t.TripID, " Database")
	}
	// GetRecords(db)

	// defer the close till after the main function has finished executing
	defer db.Close()

	fmt.Println("Database opened")
	return
}
