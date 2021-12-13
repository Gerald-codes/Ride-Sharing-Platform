package DriverDB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Declare Variables and Structure
var id string

type Driver struct { // map this type to the record in the table
	DriverID     int
	FirstName    string
	LastName     string
	MobileNo     int
	EmailAddress string
	LicenseNo    string
	Status       string
}

// Get All Driver Records
func GetRecords() (res []Driver) {
	db, Qerr := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ride_sharing")

	// handle error
	if Qerr != nil {
		panic(Qerr.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()
	results, err := db.Query("Select * FROM ride_sharing.Drivers")

	if err != nil {
		panic(err.Error())
	}

	var DriverRecords []Driver
	for results.Next() {
		// map this type to the record in the table
		var driver Driver
		err = results.Scan(&driver.DriverID, &driver.FirstName, &driver.LastName,
			&driver.MobileNo, &driver.EmailAddress, &driver.LicenseNo, &driver.Status)
		if err != nil {
			panic(err.Error())
		}
		DriverRecords = append(DriverRecords, driver)
	}
	return DriverRecords
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

	results, err := db.Query("SELECT COALESCE(MAX(DriverID),0) FROM ride_sharing.Drivers")

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

// Driver Create Account
func InsertDriver(db *sql.DB, FN string, LN string, MN int, EA string, LiN string, S string) {
	query := fmt.Sprintf("INSERT INTO ride_sharing.Drivers (FirstName, LastName, MobileNo, EmailAddress, LicenseNo, Status) VALUES ( '%s', '%s', %d,'%s','%s','%s')",
		FN, LN, MN, EA, LiN, S)
	fmt.Println("qeury", query)
	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

// Driver Edit Account
func UpdateDriver(db *sql.DB, ID int, FN string, LN string, MN int, EA string, LiN string, S string) {
	query := fmt.Sprintf(
		"UPDATE ride_sharing.Drivers SET FirstName='%s', LastName='%s',MobileNo=%d,EmailAddress='%s', LicenseNo='%s', Status='%s' WHERE DriverID=%d ",
		FN, LN, MN, EA, LiN, S, ID)

	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

// Driver Update Status
func UpdateStatus(db *sql.DB, ID int, S string) {
	query := fmt.Sprintf(
		"UPDATE ride_sharing.Driver SET Status='%s' WHERE DriverID=%d",
		S, ID)

	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

// Main
func DriverDB(method string, driver Driver) {
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ride_sharing")

	// handle error
	if err != nil {
		panic(err.Error())
	}
	if method == "Insert" {
		fmt.Println("Inserted ", driver.FirstName, " Database")
		InsertDriver(db, driver.FirstName, driver.LastName, driver.MobileNo, driver.EmailAddress, driver.LicenseNo, driver.Status)
	} else if method == "Update" {
		fmt.Println("Updated ", driver.FirstName, " Database")
		UpdateDriver(db, driver.DriverID, driver.FirstName, driver.LastName, driver.MobileNo, driver.EmailAddress, driver.LicenseNo, driver.Status)
	}

	// defer the close till after the main function has finished executing
	defer db.Close()
	fmt.Println("Driver Database opened")

}
