package REST

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type driverInfo struct {
	DriverID     int    `json:"ID"`
	FirstName    string `json:"First Name"`
	LastName     string `json:"Last Name"`
	MobileNo     int    `json:"Mobile No"`
	EmailAddress string `json:"Email Address"`
	LicenseNo    string `json:"License No"`
}

// used for storing courses on the REST API
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

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Driver REST API!")
}

func allDrivers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "List of all drivers")

	// returns the key/value pairs in the query string as a map object
	kv := r.URL.Query()

	for k, v := range kv {
		fmt.Println(k, v) // print out the key/value pair
	}

	// returns all the courses in JSON
	json.NewEncoder(w).Encode(drivers)

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
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No driver found" + "\n"))
		}
	}

	if r.Method == "DELETE" {
		if _, ok := drivers[params["driverid"]]; ok {
			delete(drivers, params["driverid"])
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("202 - driver deleted: " +
				params["driverid"] + "\n"))
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

				if newDriver.FirstName == "" &&
					newDriver.LastName == "" &&
					newDriver.MobileNo == 0 &&
					newDriver.LicenseNo == "" &&
					newDriver.EmailAddress == "" {

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
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Driver added: " +
						params["driverid"] + "\n"))
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

				if newDriver.FirstName == "" &&
					newDriver.LastName == "" &&
					newDriver.LicenseNo == "" &&
					newDriver.EmailAddress == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply driver " +
							"information " +
							"in JSON format" + "\n"))
					return
				}

				// check if course exists; add only if
				// course does not exist
				if _, ok := drivers[params["driverid"]]; !ok {
					drivers[params["driverid"]] =
						newDriver
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Driver added: " +
						params["driverid"] + "\n"))
				} else {
					// update course
					drivers[params["driverid"]] = newDriver
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - Driver updated: " +
						params["driverid"] + "\n"))
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

func REST() {

	// instantiate courses
	drivers = make(map[string]driverInfo)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/drivers", allDrivers)
	router.HandleFunc("/api/v1/drivers/{driverid}", driver).Methods(
		"GET", "PUT", "POST", "DELETE")

	fmt.Println("Listening at port 1000")
	log.Fatal(http.ListenAndServe(":1000", router))
}
