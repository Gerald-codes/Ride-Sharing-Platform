package DriverAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://localhost:1000/api/v1/drivers"
const key = "2c78afaf-97da-4816-bbee-9ad239abb296"

// Driver API Functions
func GetDriver(code string) {
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

func GetAllDriver(code string) {
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

func AddDriver(code string, jsonData map[string]interface{}) {
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(baseURL+"/"+code+"?key="+key,
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		// fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

func UpdateDriver(code string, jsonData map[string]interface{}) {
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

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

func DeleteDriver(code string) {

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

// Passenger API Functions
func AddPassenger(code string, jsonData map[string]interface{}) {
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(baseURL+"/"+code+"?key="+key,
		"application/json", bytes.NewBuffer(jsonValue))
	fmt.Println("\nApi called")
	fmt.Println("TEST", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}
