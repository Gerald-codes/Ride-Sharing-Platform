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

func GetAllTrips(code string) {
	url := baseURL
	if code != "" {
		url = baseURL + "/" + code + "?key=" + key + "&filter_by=nric"
	}
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode, "Successfully retrieved Trip History")
		fmt.Println(string(data))
		response.Body.Close()
	}
}

func GetPendingTrips() {
	url := baseURL + "?key=" + key + "&filter_by=pending"
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		// for i, v := range data {
		// 	fmt.Println("TripID\tFirstName\tLastName\tMobileNo\tPickUpPostalCode\tDropOffPostalCode\t")
		// 	fmt.Print()
		// }
		response.Body.Close()
	}
}

func AddTrip(code string, jsonData map[string]interface{}) {
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post(baseURL+"/"+code+"?key="+key,
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	// data, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode == 200 {
		fmt.Println("Trip successfully created")
	}
	response.Body.Close()

}

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

func UpdateTrip(code string, jsonData map[string]interface{}) {
	jsonValue, _ := json.Marshal(jsonData)

	request, NRerr := http.NewRequest(http.MethodPut,
		baseURL+"/"+code+"?key="+key,
		bytes.NewBuffer(jsonValue))
	fmt.Print(jsonData, "API")
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
