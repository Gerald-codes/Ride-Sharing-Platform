package PassengerAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Assignm API URL & KEY
const baseURL = "http://localhost:1001/api/v1/passengers"
const key = "2c78afaf-97da-4816-bbee-9ad239abb296"

// Passenger API Functions
// Get Passenger API (Not Used in code)
func GetPassenger(code string) {
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

// Add Passenger API
func AddPassenger(code string, jsonData map[string]interface{}) {
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(baseURL+"/"+code+"?key="+key,
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

// Update Passenger API
func UpdatePassenger(code string, jsonData map[string]interface{}) {
	jsonValue, _ := json.Marshal(jsonData)

	request, NRerr := http.NewRequest(http.MethodPut,
		baseURL+"/"+code+"?key="+key,
		bytes.NewBuffer(jsonValue))

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
		fmt.Println(string(data))
		response.Body.Close()
	}
}

// Get Latest Passenger ID API
func GetLatestPassengerID() (res string) {
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
