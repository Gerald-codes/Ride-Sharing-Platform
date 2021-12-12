package main

import (
	"encoding/json"
	"fmt"
	DriverDB "importMods/Driver_MS/Database"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type driverInfo struct {
	NRICNo       string `json:"NRICNo"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	MobileNo     int    `json:"MobileNo"`
	EmailAddress string `json:"EmailAddress"`
	LicenseNo    string `json:"LicenseNo"`
	Status       string `json:"Status"`
}

// used for storing driver on the REST API
var drivers map[string]driverInfo

func validKey(r *http.Request) bool {
	v := r.URL.Query()
	if key, ok := v["key"]; ok {
		if key[0] == "2c78afaf-97da-4816-bbee-9ad239abb296" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func allDrivers(w http.ResponseWriter, r *http.Request) {
	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key\n"))
		return
	}
	// // fmt.Fprintf(w, "List of all drivers")

	// // returns the key/value pairs in the query string as a map object
	// kv := r.URL.Query()

	// for k, v := range kv {
	// 	fmt.Println(k, v) // print out the key/value pair
	// }
	// // returns all the Drivers in JSON
	// json.NewEncoder(w).Encode(drivers)

}

func driver(w http.ResponseWriter, r *http.Request) {
	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key\n"))
		return
	}

	params := mux.Vars(r)

	if r.Method == "GET" {
		if _, ok := drivers[params["nricno"]]; ok {
			json.NewEncoder(w).Encode(
				drivers[params["nricno"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No driver found" + "\n"))
		}
	}

	if r.Method == "DELETE" {
		if _, ok := drivers[params["nricno"]]; ok {
			delete(drivers, params["nricno"])
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("202 - driver deleted: " +
				params["nricno"] + "\n"))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No driver found" + "\n"))
		}
	}

	if r.Header.Get("Content-type") == "application/json" {

		// POST is for creating new Driver
		if r.Method == "POST" {

			// read the string sent to the service
			var newDriver driverInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				// convert JSON to object
				json.Unmarshal(reqBody, &newDriver)

				if newDriver.FirstName == "" {

					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply driver " +
							"information " + "in JSON format" + "\n"))
					return
				}

				// check if driver exists; add only if driver does not exist
				if _, ok := drivers[params["nricno"]]; !ok {
					drivers[params["nricno"]] = newDriver
					AddDriverToDB(newDriver)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Driver added: " +
						params["nricno"] + "\n"))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate Driver ID" + "\n"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply driver information " +
					"in JSON format" + "\n"))
			}
		}

		//---PUT is for creating or updating
		// existing driver---
		if r.Method == "PUT" {
			var newDriver driverInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newDriver)

				if newDriver.FirstName == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply driver " +
							"information " +
							"in JSON format" + "\n"))
					return
				}

				// check if driver exists; add only if
				// driver does not exist
				if _, ok := drivers[params["nricno"]]; !ok {
					drivers[params["nricno"]] =
						newDriver
					AddDriverToDB(newDriver)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Driver added: " +
						params["nricno"] + "\n"))
				} else {
					// update driver
					drivers[params["nricno"]] = newDriver
					UpdateDriverToDB(newDriver)
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - Driver updated: " +
						params["nricno"] + "\n"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " +
					"driver information " +
					"in JSON format" + "\n"))
			}
		}
	}
}

// DB Functions
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

func GetAllDriver() {
	// DriverDB.GetRecords()
}
func GetSpecDriver() {
	// DriverDB.GetRecords()
}

func AddDriverToDB(driverInfo driverInfo) {
	fmt.Println("DRIVER: ", driverInfo)
	var driver Driver
	driver.FirstName = driverInfo.FirstName
	driver.LastName = driverInfo.LastName
	driver.NRICNo = driverInfo.NRICNo
	driver.MobileNo = driverInfo.MobileNo
	driver.LicenseNo = driverInfo.LicenseNo
	driver.Status = driverInfo.Status
	driver.EmailAddress = driverInfo.EmailAddress
	DriverDB.DriverDB("Insert", DriverDB.Driver(driver))
}

func UpdateDriverToDB(driverInfo driverInfo) {
	fmt.Println("DRIVER: ", driverInfo)
	var driver Driver
	driver.FirstName = driverInfo.FirstName
	driver.LastName = driverInfo.LastName
	driver.NRICNo = driverInfo.NRICNo
	driver.MobileNo = driverInfo.MobileNo
	driver.LicenseNo = driverInfo.LicenseNo
	driver.Status = driverInfo.Status
	driver.EmailAddress = driverInfo.EmailAddress
	DriverDB.DriverDB("Update", DriverDB.Driver(driver))
}

func main() {

	// instantiate driver
	drivers = make(map[string]driverInfo)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/drivers", allDrivers)
	router.HandleFunc("/api/v1/drivers/{nricno}", driver).Methods(
		"GET", "PUT", "POST")

	fmt.Println("Listening at port 1000")
	log.Fatal(http.ListenAndServe(":1000", router))
}
