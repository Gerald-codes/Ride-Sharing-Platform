package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Passenger struct { // map this type to the record in the table
	FirstName    string
	LastName     string
	MobileNumber int
	EmailAddress string
}

func EditRecord(db *sql.DB, ID int, FN string, LN string, EA string) {
	query := fmt.Sprintf(
		"UPDATE Passenger SET FirstName='%s', LastName='%s', EmailAddress='%s' WHERE MobileNumber=%d",
		FN, LN, EA, ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func InsertRecord(db *sql.DB, MN int, FN string, LN string, EA string) {
	query := fmt.Sprintf("INSERT INTO Passenger VALUES (%d, '%s', '%s','%s')",
		MN, FN, LN, EA)
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
		var passenger Passenger
		err = results.Scan(&passenger.MobileNumber, &passenger.FirstName,
			&passenger.LastName, &passenger.EmailAddress)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(passenger.MobileNumber, passenger.FirstName,
			passenger.LastName, passenger.EmailAddress)
	}
}

func main() {
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ride_sharing")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	GetRecords(db)

	// defer the close till after the main function has finished executing
	defer db.Close()

	fmt.Println("Database opened")

}
