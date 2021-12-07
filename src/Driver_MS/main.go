package DriverMain

import (
	"bytes"
	"encoding/json"
	"fmt"
	DriverDB "importMods/Driver_MS/Database"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://localhost:1000/api/v1/drivers"
const key = "2c78afaf-97da-4816-bbee-9ad239abb296"

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

func AddDriver(code string, jsonData map[string]interface{}) {
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(baseURL+"/"+code+"?key="+key,
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

func UpdateDriver(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)

	request, err := http.NewRequest(http.MethodPut,
		baseURL+"/"+code+"?key="+key,
		bytes.NewBuffer(jsonValue))

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

	request, err := http.NewRequest(http.MethodDelete,
		baseURL+"/"+code+"?key="+key, nil)

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
func Hello() {
	fmt.Println("SADLKS")
}

// DB Functions
type Driver struct { // map this type to the record in the table
	DriverID     int
	FirstName    string
	LastName     string
	MobileNo     int
	EmailAddress string
	LicenseNo    string
	Status       string
	ICNo         string
}

func GetAllDriver() {
	// DriverDB.GetRecords()
}

func AddDriverToDB(jsonData map[string]interface{}) {
	var driver Driver
	driver.DriverID = 0
	driver.FirstName = jsonData["First Name"].(string)
	driver.LastName = jsonData["Last Name"].(string)
	driver.ICNo = jsonData["IC No"].(string)
	driver.MobileNo = jsonData["Mobile No"].(int)
	driver.LicenseNo = jsonData["License No"].(string)
	driver.Status = jsonData["Status"].(string)
	driver.EmailAddress = jsonData["Email Address"].(string)
	DriverDB.DriverDB("Insert", DriverDB.Driver(driver))
}
