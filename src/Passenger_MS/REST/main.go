package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type passengerInfo struct {
	PassengerID  int    `json:"ID"`
	FirstName    string `json:"First Name"`
	LastName     string `json:"Last Name"`
	EmailAddress string `json:"Email Address"`
}

// used for storing courses on the REST API
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

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Passenger REST API!")
}

func allPassengers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "List of all passengers")

	// returns the key/value pairs in the query string as a map object
	kv := r.URL.Query()

	for k, v := range kv {
		fmt.Println(k, v) // print out the key/value pair
	}

	// returns all the courses in JSON
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
		if _, ok := passengers[params["passengerid"]]; ok {
			json.NewEncoder(w).Encode(
				passengers[params["passengerid"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No passenger found" + "\n"))
		}
	}

	if r.Method == "DELETE" {
		if _, ok := passengers[params["passengerid"]]; ok {
			delete(passengers, params["passengerid"])
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("202 - driver deleted: " +
				params["passengerid"] + "\n"))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No passenger found" + "\n"))
		}
	}

	if r.Header.Get("Content-type") == "application/json" {

		// POST is for creating new Driver
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

				// check if driver exists; add only if driver does not exist
				if _, ok := passengers[params["passengerid"]]; !ok {
					passengers[params["passengerid"]] = newPassenger
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
		// existing driver---
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

				// check if course exists; add only if
				// course does not exist
				if _, ok := passengers[params["passengerid"]]; !ok {
					passengers[params["passengerid"]] =
						newPassenger
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Passenger added: " +
						params["passengerid"] + "\n"))
				} else {
					// update course
					passengers[params["passengerid"]] = newPassenger
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

func main() {

	// instantiate courses
	passengers = make(map[string]passengerInfo)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)
	router.HandleFunc("/api/v1/passengers", allPassengers)
	router.HandleFunc("/api/v1/passengers/{passengerid}", passenger).Methods(
		"GET", "PUT", "POST", "DELETE")

	fmt.Println("Listening at port 1001")
	log.Fatal(http.ListenAndServe(":1001", router))
}
