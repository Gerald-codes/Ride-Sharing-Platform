package main

import (
	"encoding/json"
	"fmt"
	PassengerDB "importMods/Passenger_MS/Database"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type passengerInfo struct {
	FirstName    string `json:"First Name"`
	LastName     string `json:"Last Name"`
	MobileNo     int    `json:"Mobile No"`
	EmailAddress string `json:"Email Address"`
}

// used for storing passengers on the REST API
var passengers map[string]passengerInfo

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

func allPassengers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "List of all passengers")

	// returns the key/value pairs in the query string as a map object
	kv := r.URL.Query()

	for k, v := range kv {
		fmt.Println(k, v) // print out the key/value pair
	}

	// returns all the passengers in JSON
	json.NewEncoder(w).Encode(passengers)

}

func passenger(w http.ResponseWriter, r *http.Request) {
	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key\n"))
		return
	}

	params := mux.Vars(r)

	if r.Method == "GET" {
		if _, ok := passengers[params["moibleno"]]; ok {
			json.NewEncoder(w).Encode(
				passengers[params["moibleno"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No passenger found" + "\n"))
		}
	}

	if r.Method == "DELETE" {
		if _, ok := passengers[params["moibleno"]]; ok {
			delete(passengers, params["moibleno"])
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("202 - passenger deleted: " +
				params["moibleno"] + "\n"))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No passenger found" + "\n"))
		}
	}

	if r.Header.Get("Content-type") == "application/json" {

		// POST is for creating new Passenger
		if r.Method == "POST" {
			// read the string sent to the service
			var newPassenger passengerInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				// convert JSON to object
				json.Unmarshal(reqBody, &newPassenger)

				if newPassenger.FirstName == "" &&
					newPassenger.LastName == "" &&
					newPassenger.EmailAddress == "" {

					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply passenger " +
							"information " + "in JSON format" + "\n"))
					return
				}

				// check if passenger exists; add only if passenger does not exist
				if _, ok := passengers[params["moibleno"]]; !ok {
					passengers[params["moibleno"]] = newPassenger
					AddPassengerToDB(newPassenger)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Passenger added: " +
						params["moibleno"] + "\n"))

				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate Passenger ID" + "\n"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply passenger information " +
					"in JSON format" + "\n"))
			}
		}

		//---PUT is for creating or updating
		// existing passenger---
		if r.Method == "PUT" {
			var newPassenger passengerInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newPassenger)

				if newPassenger.FirstName == "" &&
					newPassenger.LastName == "" &&
					newPassenger.EmailAddress == "" {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply passenger " +
							"information " +
							"in JSON format" + "\n"))
					return
				}

				// check if Passenger exists; add only if
				// Passenger does not exist
				if _, ok := passengers[params["moibleno"]]; !ok {
					passengers[params["moibleno"]] =
						newPassenger
					AddPassengerToDB(newPassenger)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Passenger added: " +
						params["moibleno"] + "\n"))
				} else {
					// update Passenger
					passengers[params["moibleno"]] = newPassenger
					UpdatePassengerToDB(newPassenger)
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - Passenger updated: " +
						params["moibleno"] + "\n"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " +
					"Passenger information " +
					"in JSON format" + "\n"))
			}
		}
	}
}

// DB Functions
type Passenger struct { // map this type to the record in the table
	PassengerID  int
	FirstName    string
	LastName     string
	MobileNo     int
	EmailAddress string
}

func AddPassengerToDB(passengerInfo passengerInfo) {
	var passenger Passenger
	passenger.FirstName = passengerInfo.FirstName
	passenger.LastName = passengerInfo.LastName
	passenger.MobileNo = passengerInfo.MobileNo
	passenger.EmailAddress = passengerInfo.EmailAddress
	PassengerDB.PassengerDB("Insert", PassengerDB.Passenger(passenger))
}

func UpdatePassengerToDB(passengerInfo passengerInfo) {
	var passenger Passenger
	passenger.FirstName = passengerInfo.FirstName
	passenger.LastName = passengerInfo.LastName
	passenger.MobileNo = passengerInfo.MobileNo
	passenger.EmailAddress = passengerInfo.EmailAddress
	PassengerDB.PassengerDB("Update", PassengerDB.Passenger(passenger))
}

func main() {
	// instantiate passengers
	passengers = make(map[string]passengerInfo)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/passengers", allPassengers)
	router.HandleFunc("/api/v1/passengers/{moibleno}", passenger).Methods(
		"GET", "PUT", "POST", "DELETE")

	fmt.Println("Listening at port 1001")
	log.Fatal(http.ListenAndServe(":1001", router))
}
