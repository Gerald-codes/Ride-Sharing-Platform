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
	var User, PassengerOption, DriverOption float32
	var PassengerMenu [5]string
	PassengerMenu[0] = "\n[1] Create Account"     // Assign a value to the first element
	PassengerMenu[1] = "\n[2] Edit Account"       // Assign a value to the second element
	PassengerMenu[2] = "\n[3] Start New Trip"     // Assign a value to the third element
	PassengerMenu[3] = "\n[4] View Trips History" // Assign a value to the Fouth element
	PassengerMenu[4] = "\n[5] Back"               // Assign a value to the Fifth element
	var DriverMenu [4]string
	DriverMenu[0] = "\n[1] Create Account"     // Assign a value to the first element
	DriverMenu[1] = "\n[2] Edit Account"       // Assign a value to the second element
	DriverMenu[2] = "\n[3] View Pending Trips" // Assign a value to the third element
	DriverMenu[3] = "\n[4] Back"               // Assign a value to the Fouth element

	var UserOption [3]string
	UserOption[0] = "\n[1] Passenger" // Assign a value to the first element
	UserOption[1] = "\n[2] Driver"    // Assign a value to the second element
	UserOption[2] = "\n["
	var LN, FN, EA string
	var MN, ID int

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
					jsonData := map[string]interface{}{"FirstName": FN, "LastName": LN,
						"MobileNo": MN, "EmailAddress": EA}
					PassengerAPI.AddPassenger(PassengerAPI.GetLatestPassengerID(), jsonData)
				} else if PassengerOption == 2 {
					fmt.Println("\n----------Edit Passenger Details----------")
					fmt.Print("PassengerID: ")
					fmt.Scanln(&ID)
					fmt.Print("First Name: ")
					fmt.Scanln(&FN)
					fmt.Print("Last Name: ")
					fmt.Scanln(&LN)
					fmt.Print("Mobile Number: ")
					fmt.Scanln(&MN)
					fmt.Print("Email Address: ")
					fmt.Scanln(&EA)
					jsonData := map[string]interface{}{"PassengerID": ID, "FirstName": FN, "LastName": LN,
						"MobileNo": MN, "EmailAddress": EA}
					PassengerAPI.UpdatePassenger(strconv.Itoa(ID), jsonData)
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
						"DropOffPostalCode": DropOffPC, "Status": "Pending"}
					TripAPI.AddTrip(id, jsonData)
				} else if PassengerOption == 4 {
					fmt.Println("\n----------View Trips History----------")
					fmt.Print("PassengerID: ")
					fmt.Scanln(&ID)
					TripAPI.GetAllTrips(strconv.Itoa(ID))
				} else if PassengerOption == 5 {
					break
				} else {
					fmt.Println("\nInvalid Option")
				}
			}
		} else if User == 2 {
			for {
				fmt.Println("\n-------------Driver Menu--------------", DriverMenu[0]+DriverMenu[1]+DriverMenu[2])
				fmt.Println("\nEnter an option: ")
				fmt.Scanln(&DriverOption)
				if DriverOption == 1 {
					fmt.Println("\n----------Create Driver----------")
					var NRIC, LN, FN, EA, LiN string
					var MN int
					fmt.Println("NRIC Number: ")
					fmt.Scanln(&NRIC)
					fmt.Println("First Name: ")
					fmt.Scanln(&FN)
					fmt.Println("Last Name: ")
					fmt.Scanln(&LN)
					fmt.Println("Mobile Number: ")
					fmt.Scanln(&MN)
					fmt.Println("Email Address: ")
					fmt.Scanln(&EA)
					fmt.Println("License Number: ")
					fmt.Scanln(&LiN)
					jsonData := map[string]interface{}{"NRICNo": NRIC, "FirstName": FN, "LastName": LN,
						"MobileNo": MN, "EmailAddress": EA, "LicenseNo": LiN, "Status": "Available"}
					DriverAPI.AddDriver(NRIC, jsonData)
				} else if DriverOption == 2 {
					fmt.Println("----------Edit Driver Details----------")
					var NRIC, LN, FN, EA, LiN string
					var MN int
					fmt.Println("NRIC Number: ")
					fmt.Scanln(&NRIC)
					fmt.Println("First Name: ")
					fmt.Scanln(&FN)
					fmt.Println("Last Name: ")
					fmt.Scanln(&LN)
					fmt.Println("Mobile Number: ")
					fmt.Scanln(&MN)
					fmt.Println("Email Address: ")
					fmt.Scanln(&EA)
					fmt.Println("License Number: ")
					fmt.Scanln(&LiN)
					jsonData := map[string]interface{}{"FirstName": FN, "LastName": LN,
						"MobileNo": MN, "EmailAddress": EA, "LicenseNo": LN}
					DriverAPI.UpdateDriver(NRIC, jsonData)
				} else if DriverOption == 3 {
					fmt.Println("\n----------View Pending Trips----------")
					TripAPI.GetPendingTrips()
					fmt.Print("\nSelect a TripID or press [0] for Back: ")
					var TripOption string
					fmt.Scanln(&TripOption)
					if TripOption == "0" {
						return
					} else {
						// start trip
						var DriverID, end float32
						fmt.Print("Enter DriverID: ")
						fmt.Scanln(&DriverID)
						jsonData := map[string]interface{}{"TripID": TripOption, "DriverID": DriverID, "TripStartTime": string(time.Now().Format("01-02-2006 15:04:05")), "Status": "OnGoing"}
						TripAPI.UpdateTrip(TripOption, jsonData)
						// end trip
						fmt.Println("Reached Final Destination!", "\n[1] YES")
						fmt.Print("Enter an option: ")
						fmt.Scanln(&end)
						if end == 1 {
							jsonData := map[string]interface{}{"TripID": TripOption, "TripEndTime": string(time.Now().Format("01-02-2006 15:04:05")), "Status": "Completed"}
							TripAPI.UpdateTrip(TripOption, jsonData)
						}
					}
				} else if DriverOption == 4 {
					// fmt.Println("\n----------Auto Assign Trips----------")
					// var DriverID, end float32
					// fmt.Print("Enter DriverID: ")
					// fmt.Scanln(&DriverID)
					// TripAPI.AutoAssignTrip()
				}
			}

		}
	}

}
