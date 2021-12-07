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
	ICNo         string
}

func EditRecord(db *sql.DB, ID int, FN string, LN string, MN int, EA string, LiN string) {
	query := fmt.Sprintf(
		"UPDATE Driver SET FirstName='%s', LastName='%s', MobileNo=%d , EmailAddress='%s', LicenseNo='%s' WHERE DriverID=%d",
		FN, LN, MN, EA, LiN, ID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func UpdateStatus(db *sql.DB, ID int, S string) {
	query := fmt.Sprintf(
		"UPDATE Driver SET Status='%s' WHERE DriverID=%d",
		S, ID)

	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func InsertRecord(db *sql.DB, FN string, LN string, MN int, EA string, LiN string, S string, IC string) {
	query := fmt.Sprintf("INSERT INTO Drivers (FirstName, LastName, MobileNo, EmailAddress, LicenseNo, Status,ICNo) VALUES ( '%s', '%s', %d,'%s','%s','%s','%s')",
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
			&driver.MobileNo, &driver.EmailAddress, &driver.LicenseNo, &driver.Status, &driver.ICNo)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(driver.DriverID, driver.FirstName, driver.LastName,
			driver.MobileNo, driver.EmailAddress, driver.LicenseNo, driver.Status, driver.ICNo)
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
		fmt.Println(driver.FirstName)
		InsertRecord(db, driver.FirstName, driver.LastName, driver.MobileNo, driver.EmailAddress, driver.LicenseNo, driver.Status, driver.ICNo)
	}

	GetRecords(db)

	// defer the close till after the main function has finished executing
	defer db.Close()
	fmt.Println("Database opened")

}
