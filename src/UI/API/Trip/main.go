package TripAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://localhost:1002/api/v1/trips"
const key = "2c78afaf-97da-4816-bbee-9ad239abb296"

// Trip API Functions
// Get Trips API (Not Used in code)
func GetTrip(code string) {
	url := baseURL
	if code != "" {
		url = baseURL + "/" + code + "?key=" + key
	}
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

// Passenger Get Trip History
func GetTripHistory(PID string) {
	url := baseURL + "?key=" + key + "&filter_by=id&passenger_id=" + PID
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		if len(data) != 3 {
			fmt.Println("\n", string(data))

		} else {
			fmt.Println("\nThere is no Trip History")
		}
	}
	response.Body.Close()
}

// Get Pending Trips API
func GetPendingTrips(DID string) bool {
	url := baseURL + "?key=" + key + "&filter_by=pending&driver_id=" + DID
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return false
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		if len(data) != 3 {
			fmt.Println("\n", string(data))
			response.Body.Close()
			return true
		} else {
			response.Body.Close()
			return false
		}
	}
}

// Add Trip API
func AddTrip(code string, jsonData map[string]interface{}) {
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post(baseURL+"/"+code+"?key="+key,
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	if len(data) != 3 {
		fmt.Println("\n", string(data))
	}
	response.Body.Close()

}

// Get Latest Trip ID API
func GetLatestTripID() (res string) {
	url := baseURL + "?key=" + key + "&filter_by=latest"
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		return string(data[12 : len(data)-2])
	}
	return
}

// Update Trip API
func UpdateTrip(code string, jsonData map[string]interface{}, process string) {
	jsonValue, _ := json.Marshal(jsonData)

	request, NRerr := http.NewRequest(http.MethodPut,
		baseURL+"/"+code+"?key="+key+"&initiate="+process,
		bytes.NewBuffer(jsonValue))
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
