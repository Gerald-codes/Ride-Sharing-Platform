package PassengerDB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var id string

type Passenger struct { // map this type to the record in the table
	PassengerID  int
	FirstName    string
	LastName     string
	MobileNo     int
	EmailAddress string
}

func InsertRecord(db *sql.DB, FN string, LN string, MN int, EA string) {
	query := fmt.Sprintf("INSERT INTO ride_sharing.Passengers (FirstName, LastName, MobileNo, EmailAddress) VALUES ('%s', '%s', %d,'%s')",
		FN, LN, MN, EA)
	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func UpdatePassenger(db *sql.DB, FN string, LN string, MN int, EA string, ID int) {
	query := fmt.Sprintf(
		"UPDATE ride_sharing.Passengers SET FirstName='%s', LastName='%s',MobileNo=%d ,EmailAddress='%s' WHERE PassengerID=%d",
		FN, LN, MN, EA, ID)

	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func GetRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM ride_sharing.Passengers")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var passenger Passenger
		err = results.Scan(&passenger.PassengerID, &passenger.FirstName,
			&passenger.LastName, &passenger.MobileNo, &passenger.EmailAddress)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(passenger.PassengerID, passenger.FirstName,
			passenger.LastName, passenger.MobileNo, passenger.EmailAddress)
	}
}

func GetLatestID() (res string) {
	db, Qerr := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ride_sharing")

	// handle error
	if Qerr != nil {
		panic(Qerr.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()

	results, err := db.Query("SELECT MAX(PassengerID) FROM ride_sharing.Passengers")

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

func PassengerDB(method string, p Passenger) {
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ride_sharing")

	// handle error
	if err != nil {
		panic(err.Error())
	}
	if method == "Insert" {
		InsertRecord(db, p.FirstName, p.LastName, p.MobileNo, p.EmailAddress)
		fmt.Println("Inserted ", p.FirstName, " Database")
	} else if method == "Update" {
		UpdatePassenger(db, p.FirstName, p.LastName, p.MobileNo, p.EmailAddress, p.PassengerID)
		fmt.Println("Updated ", p.FirstName, " Database")
	}

	GetRecords(db)

	// defer the close till after the main function has finished executing
	defer db.Close()

	fmt.Println("Passenger Database opened")

}
