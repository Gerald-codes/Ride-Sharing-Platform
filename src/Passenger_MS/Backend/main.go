package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	PassengerDB "importMods/Passenger_MS/Database"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type passengerInfo struct {
	PassengerID  int    `json:"PassengerID"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	MobileNo     int    `json:"MobileNo"`
	EmailAddress string `json:"EmailAddress"`
}
type id struct {
	LatestID int `json:"LatestID"`
}

// used for storing passengers on the REST API
var passengers map[string]passengerInfo
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

func allPassengers(w http.ResponseWriter, r *http.Request) {
	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key\n"))
		return
	}
	v := r.URL.Query()
	if filter_by, ok := v["filter_by"]; ok {
		if filter_by[0] == "latest" {
			getLatest(w, r)
		}
	} else {

	}

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
				fmt.Print(&newPassenger)
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
				if _, ok := passengers[params["passengerid"]]; !ok {
					passengers[params["passengerid"]] = newPassenger
					AddPassengerToDB(newPassenger)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Passenger added: " +
						params["passengerid"] + "\n"))

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
				if _, ok := passengers[params["passengerid"]]; !ok {
					passengers[params["passengerid"]] =
						newPassenger
					AddPassengerToDB(newPassenger)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Passenger added: " +
						params["passengerid"] + "\n"))
				} else {
					// update Passenger
					passengers[params["passengerid"]] = newPassenger
					UpdatePassengerToDB(newPassenger)
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - Passenger updated: " +
						params["passengerid"] + "\n"))
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

func getLatest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if r.Method == "GET" {
		id := PassengerDB.GetLatestID()
		strVar := id
		intVar, err := strconv.Atoi(strVar)
		// handle error
		if err != nil {
			panic(err.Error())
		}

		// Backend calls PUT Request to insert latestID
		jsonData := map[string]interface{}{"LatestID": (intVar + 1)}
		jsonValue, _ := json.Marshal(jsonData)
		request, NRerr := http.NewRequest(http.MethodPut,
			"http://localhost:1001/api/v1/passengers?key=2c78afaf-97da-4816-bbee-9ad239abb296&filter_by=latest", bytes.NewBuffer(jsonValue))

		if NRerr != nil {
			fmt.Printf("New request failed with error %s\n", NRerr)
		}

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(request)
		fmt.Println("\nUpdate Passenger Api called")
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
			w.Write([]byte("404 - No PassengerID found" + "\n"))
		}
	}
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "PUT" {
			if !validKey(r) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("401 - Invalid key\n"))
				return
			}

			var newPassengerID id
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newPassengerID)
			} else {
				fmt.Print(err, "ERROR")
			}
			if _, ok := latestid[params["latestid"]]; !ok {

				latestid[params["latestid"]] = newPassengerID
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte("201 - Latest PassengerID added: " +
					params["latestid"] + "\n"))
			} else {
				// update course
				latestid[params["latestid"]] = newPassengerID
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("202 - Passenger updated: " +
					params["pasengerid"] + "\n"))
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
	passenger.PassengerID = passengerInfo.PassengerID
	passenger.FirstName = passengerInfo.FirstName
	passenger.LastName = passengerInfo.LastName
	passenger.MobileNo = passengerInfo.MobileNo
	passenger.EmailAddress = passengerInfo.EmailAddress
	PassengerDB.PassengerDB("Update", PassengerDB.Passenger(passenger))
}

func main() {
	// instantiate passengers
	passengers = make(map[string]passengerInfo)
	latestid = make(map[string]id)
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/passengers", allPassengers).Methods("GET", "PUT")
	router.HandleFunc("/api/v1/passengers/{passengerid}", passenger).Methods(
		"GET", "PUT", "POST")

	fmt.Println("Listening at port 1001")
	log.Fatal(http.ListenAndServe(":1001", router))
}
