package DriverDB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Driver struct { // map this type to the record in the table
	DriverID     int
	FirstName    string
	LastName     string
	MobileNo     int
	EmailAddress string
	LicenseNo    string
	Status       string
	NRICNo       string
}

func UpdateStatus(db *sql.DB, ID int, S string) {
	query := fmt.Sprintf(
		"UPDATE ride_sharing.Driver SET Status='%s' WHERE DriverID=%d",
		S, ID)

	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func UpdateDriver(db *sql.DB, FN string, LN string, MN int, EA string, LiN string, S string, IC string) {
	query := fmt.Sprintf(
		"UPDATE ride_sharing.Drivers SET FirstName='%s', LastName='%s',MobileNo=%d,EmailAddress='%s', LicenseNo='%s', Status='%s' WHERE NRICNO='%s'",
		FN, LN, MN, EA, LiN, S, IC)

	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func InsertDriver(db *sql.DB, FN string, LN string, MN int, EA string, LiN string, S string, IC string) {
	query := fmt.Sprintf("INSERT INTO ride_sharing.Drivers (FirstName, LastName, MobileNo, EmailAddress, LicenseNo, Status, NRICNo) VALUES ( '%s', '%s', %d,'%s','%s','%s','%s')",
		FN, LN, MN, EA, LiN, S, IC)
	fmt.Println("qeury", query)
	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func GetRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM ride_sharing.Drivers")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var driver Driver
		err = results.Scan(&driver.DriverID, &driver.FirstName, &driver.LastName,
			&driver.MobileNo, &driver.EmailAddress, &driver.LicenseNo, &driver.Status, &driver.NRICNo)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(driver.DriverID, driver.FirstName, driver.LastName,
			driver.MobileNo, driver.EmailAddress, driver.LicenseNo, driver.Status, driver.NRICNo)
	}
}

func GetSpecRecord(db *sql.DB, IC string) {
	results, err := db.Query("Select * FROM ride_sharing.Drivers WHERE NRICNo ='%s'", IC)

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var driver Driver
		err = results.Scan(&driver.DriverID, &driver.FirstName, &driver.LastName,
			&driver.MobileNo, &driver.EmailAddress, &driver.LicenseNo, &driver.Status, &driver.NRICNo)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(driver.DriverID, driver.FirstName, driver.LastName,
			driver.MobileNo, driver.EmailAddress, driver.LicenseNo, driver.Status, driver.NRICNo)
	}
}

func DriverDB(method string, driver Driver) {
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ride_sharing")

	// handle error
	if err != nil {
		panic(err.Error())
	}
	if method == "Insert" {
		fmt.Println("Inserted ", driver.FirstName, " Database")
		InsertDriver(db, driver.FirstName, driver.LastName, driver.MobileNo, driver.EmailAddress, driver.LicenseNo, driver.Status, driver.NRICNo)
	} else if method == "Update" {
		fmt.Println("Updated ", driver.FirstName, " Database")
		UpdateDriver(db, driver.FirstName, driver.LastName, driver.MobileNo, driver.EmailAddress, driver.LicenseNo, driver.Status, driver.NRICNo)
	}

	GetRecords(db)

	// defer the close till after the main function has finished executing
	defer db.Close()
	fmt.Println("Driver Database opened")

}
