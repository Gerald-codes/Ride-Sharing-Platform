package PassengerAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://localhost:1001/api/v1/passengers"
const key = "2c78afaf-97da-4816-bbee-9ad239abb296"

// Passenger API Functions
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

func GetAllPassenger(code string) {
	url := baseURL
	if code != "" {
		url = baseURL
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

func AddPassenger(code string, jsonData map[string]interface{}) {
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(baseURL+"/"+code+"?key="+key,
		"application/json", bytes.NewBuffer(jsonValue))

	fmt.Println("\nInsert Passenger Api called")

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

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
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

func DeletePassenger(code string) {

	request, NRerr := http.NewRequest(http.MethodDelete,
		baseURL+"/"+code+"?key="+key, nil)

	if NRerr != nil {
		fmt.Printf("New request failed with error %s\n", NRerr)
	}

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
}
