package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	DriverDB "importMods/Driver_MS/Database"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Create
type driverInfo struct {
	DriverID     int    `json:"DriverID"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	MobileNo     int    `json:"MobileNo"`
	EmailAddress string `json:"EmailAddress"`
	LicenseNo    string `json:"LicenseNo"`
	Status       string `json:"Status"`
}
type id struct {
	LatestID int `json:"LatestID"`
}

// Used for storing driver on the REST API
var drivers map[string]driverInfo
var latestid map[string]id

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

func CallGetAll() {
	driverArray := DriverDB.GetRecords()
	jsonValue, _ := json.Marshal(driverArray)
	request, NRerr := http.NewRequest(http.MethodPut,
		"http://localhost:1000/api/v1/drivers?key=2c78afaf-97da-4816-bbee-9ad239abb296&getall=true", bytes.NewBuffer(jsonValue))

	if NRerr != nil {
		fmt.Printf("New request failed with error %s\n", NRerr)
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

func allDrivers(w http.ResponseWriter, r *http.Request) {
	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key\n"))
		return
	}

	v := r.URL.Query()
	// Normal Call with Parameter Key only, Calls for Get All Driver records
	if len(v) == 1 {
		CallGetAll()
		enc := json.NewEncoder(w)
		enc.SetIndent("", "    ")
		enc.Encode(drivers)
	} else if getall, ok := v["getall"]; ok {
		if getall[0] == "true" {
			var newDriverDetail []driverInfo

			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newDriverDetail)
			} else {
				fmt.Print(err, "ERROR")
			}
			var id []string
			for i, v := range newDriverDetail {
				id = append(id, strconv.Itoa(v.DriverID))
				drivers[id[i]] = newDriverDetail[i]
			}
			w.WriteHeader(http.StatusAccepted)
		}
	} else if filter_by, ok := v["filter_by"]; ok {
		if filter_by[0] == "latest" {
			getLatest(w, r)
		}
	}
}

// Get Latest Driver ID
func getLatest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if r.Method == "GET" {
		// GET Latest Driver ID
		id := DriverDB.GetLatestID()
		strVar := id
		// atoi convert string to int
		intVar, err := strconv.Atoi(strVar)
		// handle error
		if err != nil {
			panic(err.Error())
		}

		// Backend calls PUT Request to insert DriverID, increase latest by 1
		jsonData := map[string]interface{}{"LatestID": (intVar + 1)}
		jsonValue, _ := json.Marshal(jsonData)
		request, NRerr := http.NewRequest(http.MethodPut,
			"http://localhost:1000/api/v1/drivers?key=2c78afaf-97da-4816-bbee-9ad239abb296&filter_by=latest", bytes.NewBuffer(jsonValue))

		if NRerr != nil {
			fmt.Printf("New request failed with error %s\n", NRerr)
		}

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(request)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(response.StatusCode)
			fmt.Println(string(data))
			response.Body.Close()
		}
		if _, ok := latestid[params["latestid"]]; ok {
			json.NewEncoder(w).Encode(
				latestid[params["latestid"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No DriverID found" + "\n"))
		}
	}
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "PUT" {
			if !validKey(r) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("401 - Invalid key\n"))
				return
			}

			var newDriverID id
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newDriverID)
			} else {
				fmt.Print(err, "ERROR")
			}
			latestid[params["latestid"]] = newDriverID
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("201 - Latest DriverID added: " +
				params["latestid"] + "\n"))
		}
	}
}

func driver(w http.ResponseWriter, r *http.Request) {
	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key\n"))
		return
	}

	params := mux.Vars(r)

	if r.Method == "GET" {
		if _, ok := drivers[params["driverid"]]; ok {
			json.NewEncoder(w).Encode(
				drivers[params["driverid"]])
		} else { // Scenario where GetAll Drivers wasnt called before
			CallGetAll()
			if _, ok := drivers[params["driverid"]]; ok {
				json.NewEncoder(w).Encode(
					drivers[params["driverid"]])
			} else { //Really No Driver found
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - No Driver found"))
			}
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

				if newDriver.DriverID == 0 {

					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply driver " +
							"information " + "in JSON format" + "\n"))
					return
				}

				// check if driver exists; add only if driver does not exist
				if _, ok := drivers[params["driverid"]]; !ok {
					drivers[params["driverid"]] = newDriver
					AddDriverToDB(newDriver)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Driver added: " +
						params["driverid"] + "\n"))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate Driver ID" + "\n"))
				}
			}
		}

		//---PUT is for creating or updating
		// existing driver---
		if r.Method == "PUT" {
			var newDriver driverInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newDriver)

				if newDriver.DriverID == 0 {
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
				if _, ok := drivers[params["driverid"]]; !ok {
					drivers[params["driverid"]] =
						newDriver
					AddDriverToDB(newDriver)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Driver added: " +
						params["driverid"] + "\n"))
				} else {
					// update driver
					drivers[params["driverid"]] = newDriver
					UpdateDriverToDB(newDriver)
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("\n202 - Driver updated: " +
						params["driverid"] + "\n"))
				}
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
}

func AddDriverToDB(driverInfo driverInfo) {
	fmt.Println("DRIVER: ", driverInfo)
	var driver Driver
	driver.FirstName = driverInfo.FirstName
	driver.LastName = driverInfo.LastName
	driver.MobileNo = driverInfo.MobileNo
	driver.LicenseNo = driverInfo.LicenseNo
	driver.Status = driverInfo.Status
	driver.EmailAddress = driverInfo.EmailAddress
	DriverDB.DriverDB("Insert", DriverDB.Driver(driver))
}

func UpdateDriverToDB(driverInfo driverInfo) {
	fmt.Println("DRIVER: ", driverInfo)
	var driver Driver
	driver.DriverID = driverInfo.DriverID
	driver.FirstName = driverInfo.FirstName
	driver.LastName = driverInfo.LastName
	driver.MobileNo = driverInfo.MobileNo
	driver.LicenseNo = driverInfo.LicenseNo
	driver.Status = driverInfo.Status
	driver.EmailAddress = driverInfo.EmailAddress
	DriverDB.DriverDB("Update", DriverDB.Driver(driver))
}

func main() {

	// instantiate driver
	drivers = make(map[string]driverInfo)
	latestid = make(map[string]id)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/drivers", allDrivers)
	router.HandleFunc("/api/v1/drivers/{driverid}", driver).Methods(
		"GET", "PUT", "POST")

	fmt.Println("Listening at port 1000")
	log.Fatal(http.ListenAndServe(":1000", router))
}
