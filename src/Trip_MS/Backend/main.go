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

// used for storing courses on the REST API
var trips map[string]tripInfo
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

func getLatest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if r.Method == "GET" {
		id := TripDB.GetLatestID()
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
			"http://localhost:1002/api/v1/trips/latest?key=2c78afaf-97da-4816-bbee-9ad239abb296", bytes.NewBuffer(jsonValue))

		if NRerr != nil {
			fmt.Printf("New request failed with error %s\n", NRerr)
		}

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(request)
		fmt.Println("\nUpdate Trip Api called")
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
		fmt.Print("TESTEST")
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
				// update course
				latestid[params["latestid"]] = newTripID
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("202 - Trip updated: " +
					params["tripid"] + "\n"))
			}
		}
	}
}

func allTrips(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "List of all Trips")

	// returns the key/value pairs in the query string as a map object
	kv := r.URL.Query()

	for k, v := range kv {
		fmt.Println(k, v) // print out the key/value pair
	}

	// returns all the courses in JSON
	json.NewEncoder(w).Encode(trips)

}

func trip(w http.ResponseWriter, r *http.Request) {
	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key\n"))
		return
	}

	params := mux.Vars(r)

	if r.Method == "GET" {
		if _, ok := trips[params["tripid"]]; ok {
			json.NewEncoder(w).Encode(
				trips[params["tripid"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No Trip found" + "\n"))
		}
	}

	if r.Method == "DELETE" {
		if _, ok := trips[params["tripid"]]; ok {
			delete(trips, params["tripid"])
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("202 - Trip deleted: " +
				params["tripid"] + "\n"))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No Trip found" + "\n"))
		}
	}

	if r.Header.Get("Content-type") == "application/json" {

		// POST is for creating new Driver
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

					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply passenger " +
							"information " + "in JSON format" + "\n"))
					return
				}

				// check if driver exists; add only if driver does not exist
				if _, ok := trips[params["tripid"]]; !ok {
					trips[params["tripid"]] = newTrip
					w.WriteHeader(http.StatusCreated)
					AddTripToDB(newTrip)
					w.Write([]byte("201 - Trip added: " +
						params["tripid"] + "\n"))
				} else {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate Trip ID" + "\n"))
				}

			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply trip information " +
					"in JSON format" + "\n"))
			}
		}

		//---PUT is for creating or updating
		// existing driver---
		if r.Method == "PUT" {
			var newTrip tripInfo
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &newTrip)

				if newTrip.TripID == 0 &&
					newTrip.DropOffPostalCode == 0 &&
					newTrip.PickUpPostalCode == 0 {
					w.WriteHeader(
						http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply trip " +
							"information " +
							"in JSON format" + "\n"))
					return
				}

				// check if course exists; add only if
				// course does not exist
				if _, ok := trips[params["tripid"]]; !ok {
					trips[params["tripid"]] =
						newTrip
					w.WriteHeader(http.StatusCreated)
					AddTripToDB(newTrip)
					w.Write([]byte("201 - Trip added: " +
						params["tripid"] + "\n"))
				} else {
					// update course
					trips[params["tripid"]] = newTrip
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - Trip updated: " +
						params["tripid"] + "\n"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " +
					"Trip information " +
					"in JSON format" + "\n"))
			}
		}
	}
}

// DB Functions
type Trip struct { // map this type to the record in the table
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
	fmt.Println("HEYHEYHEY")
	var trip Trip
	trip.TripID = tripInfo.TripID
	trip.PassengerID = tripInfo.PassengerID
	trip.PickUpPostalCode = tripInfo.PickUpPostalCode
	trip.DropOffPostalCode = tripInfo.DropOffPostalCode
	trip.TripStartTime = tripInfo.TripStartTime
	trip.Status = tripInfo.Status
	TripDB.TripDB("Insert", TripDB.Trip(trip))
}

func main() {

	// instantiate courses
	trips = make(map[string]tripInfo)
	latestid = make(map[string]id)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/trips/latest", getLatest).Methods(
		"GET", "PUT")
	router.HandleFunc("/api/v1/trips", allTrips)
	router.HandleFunc("/api/v1/trips/{tripid}", trip).Methods(
		"GET", "PUT", "POST", "DELETE")

	fmt.Println("Listening at port 1002")
	log.Fatal(http.ListenAndServe(":1002", router))
}
