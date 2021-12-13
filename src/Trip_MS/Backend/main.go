package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	TripDB "importMods/Trip_MS/Database"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Declare Stuctures of objects and Variables
type tripInfo struct {
	TripID            int    `json:"TripID"`
	PassengerID       int    `json:"PassengerID"`
	DriverID          int    `json:"DriverID"`
	PickUpPostalCode  int    `json:"PickUpPostalCode"`
	DropOffPostalCode int    `json:"DropOffPostalCode"`
	TripStartTime     string `json:"TripStartTime"`
	TripEndTime       string `json:"TripEndTime"`
	Status            string `json:"Status"`
}
type id struct {
	LatestID int `json:"LatestID"`
}
type TripDetails struct { // map this type to the record in the table
	TripID            int    `json:"TripID"`
	FirstName         string `json:"FirstName"`
	LastName          string `json:"LastName"`
	MobileNo          int    `json:"MobileNo"`
	PickUpPostalCode  int    `json:"PickUpPostalCode"`
	DropOffPostalCode int    `json:"DropOffPostalCode"`
}

// used for storing Trip on the REST API
var trips map[string]tripInfo
var latestid map[string]id
var tripDetails map[string]TripDetails

// Check if parameter key is valid
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

// Get Latest ID
func getLatest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if r.Method == "GET" {
		// GET Latest Trip ID
		id := TripDB.GetLatestID()
		strVar := id
		// Atoi convert string to int
		intVar, err := strconv.Atoi(strVar)
		// handle error
		if err != nil {
			panic(err.Error())
		}

		// Backend calls PUT Request to insert latestID, increase latest by 1
		jsonData := map[string]interface{}{"LatestID": (intVar + 1)}
		jsonValue, _ := json.Marshal(jsonData)
		request, NRerr := http.NewRequest(http.MethodPut,
			"http://localhost:1002/api/v1/trips?key=2c78afaf-97da-4816-bbee-9ad239abb296&filter_by=latest", bytes.NewBuffer(jsonValue))

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
			w.Write([]byte("404 - No TripID found" + "\n"))
		}
	}
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "PUT" {
			if !validKey(r) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("401 - Invalid key\n"))
				return
			}

			var newTripID id
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newTripID)
			} else {
				fmt.Print(err, "ERROR")
			}
			if _, ok := latestid[params["latestid"]]; !ok {

				latestid[params["latestid"]] = newTripID
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte("201 - Latest TripID added: " +
					params["latestid"] + "\n"))
			} else {
				// update Trip
				latestid[params["latestid"]] = newTripID
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("202 - Trip updated: " +
					params["tripid"] + "\n"))
			}
		}
	}
}

// Get All Trip Records
func CallGetAll() {
	tripArray := TripDB.GetRecords()
	jsonValue, _ := json.Marshal(tripArray)
	request, NRerr := http.NewRequest(http.MethodPut,
		"http://localhost:1002/api/v1/trips?key=2c78afaf-97da-4816-bbee-9ad239abb296&getall=true", bytes.NewBuffer(jsonValue))

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

func allTrips(w http.ResponseWriter, r *http.Request) {
	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key\n"))
		return
	}
	v := r.URL.Query()
	// Normal Call with Parameter Key only, Calls for Get All trip records
	if len(v) == 1 {
		CallGetAll()
		enc := json.NewEncoder(w)
		enc.SetIndent("", "    ")
		enc.Encode(trips)
	} else if getall, ok := v["getall"]; ok {
		if getall[0] == "true" {
			var newTripDetail []tripInfo

			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newTripDetail)
			} else {
				fmt.Print(err, "ERROR")
			}
			var id []string
			for i, v := range newTripDetail {
				id = append(id, strconv.Itoa(v.TripID))
				trips[id[i]] = newTripDetail[i]
			}
			w.WriteHeader(http.StatusAccepted)
		}
	} else if filter_by, ok := v["filter_by"]; ok {
		if filter_by[0] == "latest" {
			getLatest(w, r)
		} else if filter_by[0] == "pending" {
			if driver_id, ok := v["driver_id"]; ok {
				getPendingTrips(w, r, driver_id[0])
			}
		} else if filter_by[0] == "id" {
			if passenger_id, ok := v["passenger_id"]; ok {
				getTripHistory(w, r, passenger_id[0])
			}
		}
	}
}

// Driver Getting pending Trips that system has assigned to them
func getPendingTrips(w http.ResponseWriter, r *http.Request, DID string) {
	var newTripDetail []TripDetails
	if r.Method == "GET" {
		intDID, _ := strconv.Atoi(DID)
		tripArray := TripDB.GetPendingRecords(intDID)
		jsonValue, _ := json.Marshal(tripArray)
		request, NRerr := http.NewRequest(http.MethodPut,
			"http://localhost:1002/api/v1/trips?key=2c78afaf-97da-4816-bbee-9ad239abb296&filter_by=pending&driver_id="+DID, bytes.NewBuffer(jsonValue))

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
		enc := json.NewEncoder(w)
		enc.SetIndent("", "    ")
		enc.Encode(tripDetails)

	}
	if r.Method == "PUT" {
		if !validKey(r) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("401 - Invalid key\n"))
			return
		}

		reqBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			json.Unmarshal(reqBody, &newTripDetail)
		} else {
			fmt.Print(err, "ERROR")
		}
		var id []string
		for i, v := range newTripDetail {
			id = append(id, strconv.Itoa(v.TripID))
			tripDetails[id[i]] = newTripDetail[i]
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("201 - Successfully GET Pending Trip\n"))
		return
	}

}

// Passenger Calling to get Trip History
func getTripHistory(w http.ResponseWriter, r *http.Request, PID string) {
	var newTripDetail []tripInfo

	if r.Method == "GET" {
		intPID, _ := strconv.Atoi(PID)
		tripArray := TripDB.GetTripHistory(intPID)
		jsonValue, _ := json.Marshal(tripArray)
		request, NRerr := http.NewRequest(http.MethodPut,
			"http://localhost:1002/api/v1/trips?key=2c78afaf-97da-4816-bbee-9ad239abb296&filter_by=id&passenger_id="+PID, bytes.NewBuffer(jsonValue))

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
		enc := json.NewEncoder(w)
		enc.SetIndent("", "    ")
		enc.Encode(trips)

	}
	if r.Method == "PUT" {
		if !validKey(r) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("401 - Invalid key\n"))
			return
		}

		reqBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			json.Unmarshal(reqBody, &newTripDetail)
		} else {
			fmt.Print(err, "ERROR")
		}
		var id []string
		for i, v := range newTripDetail {
			id = append(id, strconv.Itoa(v.PassengerID))
			trips[id[i]] = newTripDetail[i]
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("201 - Successfully GET Trip History\n"))
	}

}

// Specific Trip
func trip(w http.ResponseWriter, r *http.Request) {
	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key\n"))
		return
	}
	// Get url params
	v := r.URL.Query()
	// Get Parameters
	params := mux.Vars(r)
	// Get Specific Trip
	if r.Method == "GET" {
		if _, ok := trips[params["tripid"]]; ok {
			json.NewEncoder(w).Encode(
				trips[params["tripid"]])
		} else { // Scenario where GetAll Trips wasnt called before
			CallGetAll()
			if _, ok := trips[params["tripid"]]; ok {
				json.NewEncoder(w).Encode(
					trips[params["tripid"]])
			} else { //Really No Trip found
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - No trip found"))
				return
			}
		}
	}

	if r.Header.Get("Content-type") == "application/json" {

		// POST is for creating new Trip
		if r.Method == "POST" {

			// read the string sent to the service
			var newTrip tripInfo
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				// convert JSON to object
				json.Unmarshal(reqBody, &newTrip)
				if newTrip.TripID == 0 &&
					newTrip.DropOffPostalCode == 0 &&
					newTrip.PickUpPostalCode == 0 {

					w.WriteHeader(http.StatusUnprocessableEntity)
				}

				// check if Trip exists; add only if Trip does not exist
				if _, ok := trips[params["tripid"]]; !ok {
					trips[params["tripid"]] = newTrip
					w.WriteHeader(http.StatusCreated)
					AddTripToDB(newTrip)
					w.Write([]byte("201 - Trip added: " +
						params["tripid"]))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate Trip ID" + "\n"))
				}

			}
		}

		//---PUT is for creating or updating
		// existing Trip---
		if r.Method == "PUT" {
			var newTrip tripInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newTrip)
				if newTrip.TripID == 0 {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply Trip " +
							"information " +
							"in JSON format" + "\n"))
					return
				}
				// Update Trip
				// Check if there is initiate params
				if initiate, ok := v["initiate"]; ok {
					// if initiate is start
					if initiate[0] == "start" {
						StartTripToDB(newTrip)
						w.WriteHeader(http.StatusAccepted)
					} else if initiate[0] == "end" { // if initiate is End
						EndTripToDB(newTrip)
						w.WriteHeader(http.StatusAccepted)
					} else {
						fmt.Print("It should not reach here")
					}
				}

			}
		}
	}
}

// DB Functions
type TripDetail struct { // map this type to the record in the table
	TripID            int
	PassengerID       int
	DriverID          int
	PickUpPostalCode  int
	DropOffPostalCode int
	TripStartTime     string
	TripEndTime       string
	Status            string
}

func AddTripToDB(tripInfo tripInfo) {
	var trip TripDetail
	trip.TripID = tripInfo.TripID
	trip.PassengerID = tripInfo.PassengerID
	trip.PickUpPostalCode = tripInfo.PickUpPostalCode
	trip.DropOffPostalCode = tripInfo.DropOffPostalCode
	trip.Status = tripInfo.Status
	TripDB.TripDB("Insert", TripDB.Trip(trip))
}

func StartTripToDB(tripInfo tripInfo) {
	var trip TripDetail
	trip.TripID = tripInfo.TripID
	trip.TripStartTime = tripInfo.TripStartTime
	trip.Status = tripInfo.Status
	fmt.Print("FMT", trip)
	TripDB.TripDB("Start", TripDB.Trip(trip))

}

func EndTripToDB(tripInfo tripInfo) {
	var trip TripDetail
	trip.TripID = tripInfo.TripID
	trip.TripEndTime = tripInfo.TripEndTime
	trip.Status = tripInfo.Status
	fmt.Print("FMT", trip)
	TripDB.TripDB("End", TripDB.Trip(trip))
}

func main() {
	// instantiate trips, latetid and tripDetails
	trips = make(map[string]tripInfo)
	latestid = make(map[string]id)
	tripDetails = make(map[string]TripDetails)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/trips", allTrips).Methods("GET", "PUT")
	router.HandleFunc("/api/v1/trips/{tripid}", trip).Methods(
		"GET", "PUT", "POST")

	fmt.Println("Listening at port 1002")
	log.Fatal(http.ListenAndServe(":1002", router))
}
