package main

import (
	"fmt"
	DriverAPI "importMods/UI/API/Driver"
	PassengerAPI "importMods/UI/API/Passenger"
	TripAPI "importMods/UI/API/Trip"
	"strconv"
	"time"
)

func main() {
	var User, PassengerOption, DriverOption, TripOption float32
	var PassengerMenu [5]string
	PassengerMenu[0] = "\n[1] Create Account"     // Assign a value to the first element
	PassengerMenu[1] = "\n[2] Edit Account"       // Assign a value to the second element
	PassengerMenu[2] = "\n[3] Start New Trip"     // Assign a value to the third element
	PassengerMenu[3] = "\n[4] View Trips History" // Assign a value to the Fouth element
	PassengerMenu[4] = "\n[5] Back"               // Assign a value to the Fifth element
	var DriverMenu [3]string
	DriverMenu[0] = "\n[1] Create Account"       // Assign a value to the first element
	DriverMenu[1] = "\n[2] Edit Account"         // Assign a value to the second element
	DriverMenu[2] = "\n[3] View Available Trips" // Assign a value to the third element
	var UserOption [3]string
	UserOption[0] = "\n[1] Passenger" // Assign a value to the first element
	UserOption[1] = "\n[2] Driver"    // Assign a value to the second element
	UserOption[2] = "\n["
	var LN, FN, EA string
	var MN int

	for {
		fmt.Println("\n---------Welcome to Ride Sharing---------\n", UserOption[0]+UserOption[1])
		fmt.Print("\nEnter an option: ")
		fmt.Scanln(&User)
		if User == 1 {
			for {
				fmt.Println("\n-------------Passenger Menu--------------\n", PassengerMenu[0]+PassengerMenu[1]+PassengerMenu[2]+PassengerMenu[3]+PassengerMenu[4])
				fmt.Print("\nEnter an option: ")
				fmt.Scanln(&PassengerOption)
				if PassengerOption == 1 {
					fmt.Println("\n----------Create Passenger----------")
					fmt.Print("First Name: ")
					fmt.Scanln(&FN)
					fmt.Print("Last Name: ")
					fmt.Scanln(&LN)
					fmt.Print("Mobile Number: ")
					fmt.Scanln(&MN)
					fmt.Print("Email Address: ")
					fmt.Scanln(&EA)
					jsonData := map[string]interface{}{"First Name": FN, "Last Name": LN,
						"Mobile No": MN, "Email Address": EA}
					// jsonData := map[string]interface{}{"First Name": "Troll", "Last Name": "Ta1n",
					// 	"Mobile No": 91234567, "Email Address": "Testing@test.com"}
					// PassengerAPI.AddPassenger("91234567", jsonData)
					PassengerAPI.AddPassenger(strconv.Itoa(MN), jsonData)
				} else if PassengerOption == 2 {
					fmt.Println("\n----------Edit Passenger Details----------")
					fmt.Print("First Name: ")
					fmt.Scanln(&FN)
					fmt.Print("Last Name: ")
					fmt.Scanln(&LN)
					fmt.Print("Mobile Number: ")
					fmt.Scanln(&MN)
					fmt.Print("Email Address: ")
					fmt.Scanln(&EA)
					jsonData := map[string]interface{}{"First Name": FN, "Last Name": LN,
						"Mobile No": MN, "Email Address": EA}
					// jsonData := map[string]interface{}{"First Name": "GERALD", "Last Name": "Ta1n",
					// 	"Mobile No": 91234567, "Email Address": "Testing@test.com"}
					// PassengerAPI.AddPassenger("91234567", jsonData)
					PassengerAPI.UpdatePassenger(strconv.Itoa(MN), jsonData)
				} else if PassengerOption == 3 {
					fmt.Print("\n----------Start New Trip----------\n")
					id := TripAPI.GetLatestTripID()
					var PickUpPC, DropOffPC, PassengerID int
					fmt.Print("Enter PassengerID: ")
					fmt.Scanln(&PassengerID)
					fmt.Print("Enter Pick Up Postal Code: ")
					fmt.Scanln(&PickUpPC)
					fmt.Print("Enter Drop Off Postal Code: ")
					fmt.Scanln(&DropOffPC)
					jsonData := map[string]interface{}{"TripID": id, "PassengerID": PassengerID, "PickUpPostalCode": PickUpPC,
						"DropOffPostalCode": DropOffPC, "TripStartTime": string(time.Now().Format("01-02-2006 15:04:05")), "Status": "Pending"}
					TripAPI.AddTrip(id, jsonData)
				} else if PassengerOption == 4 {
					fmt.Println("\n----------View Trips History----------")
					fmt.Println(time.Now().Format("01-02-2006 15:04:05"))
				} else if PassengerOption == 5 {
					break
				} else {
					fmt.Println("\nInvalid Option")
				}
			}
		} else if User == 2 {
			for {
				fmt.Println("\n-------------Driver Menu--------------", DriverMenu[0]+DriverMenu[1]+DriverMenu[2])
				fmt.Println("\nEnter an ption: ")
				fmt.Scanln(&DriverOption)
				if DriverOption == 1 {
					fmt.Println("\n----------Create Driver----------")
					// var NRIC, LN, FN, EA, LiN string
					// var MN int
					// fmt.Println("NRIC Number: ")
					// fmt.Scanln(&NRIC)
					// fmt.Println("First Name: ")
					// fmt.Scanln(&FN)
					// fmt.Println("Last Name: ")
					// fmt.Scanln(&LN)
					// fmt.Println("Mobile Number: ")
					// fmt.Scanln(&MN)
					// fmt.Println("Email Address: ")
					// fmt.Scanln(&EA)
					// fmt.Println("License Number: ")
					// fmt.Scanln(&LiN)
					// jsonData := map[string]interface{}{"IC No": NRIC, "First Name": FN, "Last Name": LN,
					// 	"Mobile No": MN, "Email Address": EA, "License No": LiN, "Status": "Available"}
					jsonData := map[string]interface{}{"NRIC No": "T0112123R", "First Name": "Troll", "Last Name": "Ta1n",
						"Mobile No": 91234567, "Email Address": "Testing@test.com", "License No": "LKS123L", "Status": "Free"}
					DriverAPI.AddDriver("T0112123R", jsonData)
				} else if DriverOption == 2 {
					fmt.Println("----------Edit Driver Details----------")
					// var NRIC string
					// fmt.Println("NRIC Number: ")
					// fmt.Scanln(&NRIC)
					jsonData := map[string]interface{}{"NRIC No": "T0112123R", "First Name": "Gerald", "Last Name": "Tan",
						"Mobile No": 91234567, "Email Address": "Testing@test.com", "License No": "LKS123L", "Status": "Free"}
					DriverAPI.UpdateDriver("T0112123R", jsonData)
				} else if DriverOption == 3 {
					fmt.Println("\n----------View Available Trips----------")
					// Print Available Trips

					fmt.Println("\nEnter TripID to select Trip\n[0] Back")
					fmt.Scanln(&TripOption)
					if TripOption == 0 {
						return
					} else {
						// start trip
						// end trip
					}
				}
			}

		}
	}

}
