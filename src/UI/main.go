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

	// Set User Input variables for menu options
	var User, PassengerOption, DriverOption, end float32
	var LN, FN, EA, LiN string
	var MN, ID, PickUpPC, DropOffPC, PassengerID, DriverID int
	// Menu to select User type
	var UserOption [2]string
	UserOption[0] = "\n[1] Passenger" // Assign a value to the first element
	UserOption[1] = "\n[2] Driver"    // Assign a value to the second element

	// Passenger Main Menu
	var PassengerMenu [5]string
	PassengerMenu[0] = "\n[1] Create Account"     // Assign a value to the first element
	PassengerMenu[1] = "\n[2] Edit Account"       // Assign a value to the second element
	PassengerMenu[2] = "\n[3] Start New Trip"     // Assign a value to the third element
	PassengerMenu[3] = "\n[4] View Trips History" // Assign a value to the Fouth element
	PassengerMenu[4] = "\n[5] Back"               // Assign a value to the Fifth element

	// Driver Mian Menu
	var DriverMenu [4]string
	DriverMenu[0] = "\n[1] Create Account"     // Assign a value to the first element
	DriverMenu[1] = "\n[2] Edit Account"       // Assign a value to the second element
	DriverMenu[2] = "\n[3] View Pending Trips" // Assign a value to the third element
	DriverMenu[3] = "\n[4] Back"               // Assign a value to the Fouth element

	// For loop for the whole program so that its smooth
	for {

		// Display User Main and prompt for user input
		fmt.Println("\n---------Welcome to Ride Sharing---------\n", UserOption[0]+UserOption[1])
		fmt.Print("\nEnter an option: ")
		fmt.Scanln(&User)

		// Display Passenger Menu
		if User == 1 {
			for {
				fmt.Println("\n-------------Passenger Menu--------------\n", PassengerMenu[0]+PassengerMenu[1]+PassengerMenu[2]+PassengerMenu[3]+PassengerMenu[4])
				fmt.Print("\nEnter an option: ")
				fmt.Scanln(&PassengerOption)

				// Passenger Option 1 - Create Account
				if PassengerOption == 1 {
					fmt.Println("\n----------Create Passenger Account----------")
					fmt.Print("First Name: ")
					fmt.Scanln(&FN)
					fmt.Print("Last Name: ")
					fmt.Scanln(&LN)
					fmt.Print("Mobile Number: ")
					fmt.Scanln(&MN)
					fmt.Print("Email Address: ")
					fmt.Scanln(&EA)
					// Map Passenger Details into string:interface as MobileNumber is int
					jsonData := map[string]interface{}{"FirstName": FN, "LastName": LN,
						"MobileNo": MN, "EmailAddress": EA}
					// Call AddPassenger Function in PassengerAPI to insert Passenger into DateBase
					// Call GetLatestPassengerID in passengerAPI for latest PassengerID and use it as the code for reference
					PassengerAPI.AddPassenger(PassengerAPI.GetLatestPassengerID(), jsonData)
				} else if PassengerOption == 2 {
					// Passenger Option 2 - Edit Account
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
					// Map Passenger Details into string:interface as MobileNumber is int
					jsonData := map[string]interface{}{"PassengerID": ID, "FirstName": FN, "LastName": LN,
						"MobileNo": MN, "EmailAddress": EA}
					// Call UpdatePassenger Function in PassengerAPI to Updaet Passenger in DateBase
					// Call strconv.Itoa() to convert int to string for the code as reference
					PassengerAPI.UpdatePassenger(strconv.Itoa(ID), jsonData)
				} else if PassengerOption == 3 {
					// Passenger Option 3 - Start New Trip
					fmt.Print("\n----------Start New Trip----------\n")
					// Get Latest Trip ID for code reference
					id := TripAPI.GetLatestTripID()
					// Prompt for PassengerID as Login service was not Implemented
					fmt.Print("Enter PassengerID: ")
					fmt.Scanln(&PassengerID)
					fmt.Print("Enter Pick Up Postal Code: ")
					fmt.Scanln(&PickUpPC)
					fmt.Print("Enter Drop Off Postal Code: ")
					fmt.Scanln(&DropOffPC)
					// Convert TripID from String to Int
					intID, _ := strconv.Atoi(id)
					// Map Trip Details into string:interface as IDs are integers
					jsonData := map[string]interface{}{"TripID": intID, "PassengerID": PassengerID, "PickUpPostalCode": PickUpPC,
						"DropOffPostalCode": DropOffPC, "Status": "Pending"}
					// Call Add Trip function from TripAPI
					TripAPI.AddTrip(id, jsonData)
				} else if PassengerOption == 4 {
					// Passenger Option 4 - View Trips History
					fmt.Println("\n----------View Trips History----------")
					fmt.Print("PassengerID: ")
					fmt.Scanln(&ID)
					// Convert int ID to string
					// Call GetAllTrip to display Trip history
					TripAPI.GetTripHistory(strconv.Itoa(ID))
				} else if PassengerOption == 5 {
					// Passenger Option 5 - Back
					break
				} else {
					// Passenger Menu Invalid Option
					fmt.Println("\nInvalid Option")
				}
			}
		} else if User == 2 {
			// Diaplay Driver Menu
			for {
				fmt.Println("\n-------------Driver Menu--------------\n", DriverMenu[0]+DriverMenu[1]+DriverMenu[2]+DriverMenu[3])
				fmt.Print("\nEnter an option: ")
				fmt.Scanln(&DriverOption)
				if DriverOption == 1 {
					// Driver Option 1 - Create Driver Account
					fmt.Println("\n----------Create Driver Account----------")
					// Prompt User for Driver Details
					id := DriverAPI.GetLatestID()
					fmt.Print("First Name: ")
					fmt.Scanln(&FN)
					fmt.Print("Last Name: ")
					fmt.Scanln(&LN)
					fmt.Print("Mobile Number: ")
					fmt.Scanln(&MN)
					fmt.Print("Email Address: ")
					fmt.Scanln(&EA)
					fmt.Print("License Number: ")
					fmt.Scanln(&LiN)
					intDriverID, _ := strconv.Atoi(id)
					// Map Driver Details into string:interface as ID is int
					jsonData := map[string]interface{}{"DriverID": intDriverID, "FirstName": FN, "LastName": LN,
						"MobileNo": MN, "EmailAddress": EA, "LicenseNo": LiN, "Status": "Available"}
					// Call AddDriver in DriverAPI to Insert Driver to DataBase
					DriverAPI.AddDriver(id, jsonData)
				} else if DriverOption == 2 {
					// Driver Option 2 - Edit Driver Details
					fmt.Println("----------Edit Driver Details----------")
					// Prompt User for Driver Details
					// Prompt for DriverID as Login service was not Implemented
					fmt.Print("Driver ID: ")
					fmt.Scanln(&DriverID)
					fmt.Print("First Name: ")
					fmt.Scanln(&FN)
					fmt.Print("Last Name: ")
					fmt.Scanln(&LN)
					fmt.Print("Mobile Number: ")
					fmt.Scanln(&MN)
					fmt.Print("Email Address: ")
					fmt.Scanln(&EA)
					fmt.Print("License Number: ")
					fmt.Scanln(&LiN)
					// Map Driver Details into string:interface as DriverID is integers
					jsonData := map[string]interface{}{"DriverID": DriverID, "FirstName": FN, "LastName": LN,
						"MobileNo": MN, "EmailAddress": EA, "LicenseNo": LN, "Status": "Available"}
					// Call UpdateDriver in DriverAPI to Update Driver to DataBase
					DriverAPI.UpdateDriver(strconv.Itoa(DriverID), jsonData)
				} else if DriverOption == 3 {
					// Driver Option 3 - View Pending Trips
					for {
						fmt.Println("\n----------View Pending Trips----------")
						fmt.Print("Enter DriverID: ")
						fmt.Scanln(&DriverID)
						pTrip := TripAPI.GetPendingTrips(strconv.Itoa(DriverID))
						if pTrip {
							fmt.Print("\nEnter TripID or press [0] for Back: ")
							var TripOption string
							fmt.Scanln(&TripOption)
							if TripOption == "0" {
								// Back Option
								break
							} else {
								// Initiative Start Trip
								// Map Driver Details into string:interface as TripID is int
								intTrip, _ := strconv.Atoi(TripOption)
								jsonData := map[string]interface{}{"TripID": intTrip, "TripStartTime": string(time.Now().Format("01-02-2006 15:04:05")), "Status": "OnGoing"}
								TripAPI.UpdateTrip(TripOption, jsonData, "start")
								// Update Driver Status to Busy
								// Map Driver Details into string:interface as DriverID is integers
								jsonDataS := map[string]interface{}{"DriverID": DriverID, "FirstName": FN, "LastName": LN,
									"MobileNo": MN, "EmailAddress": EA, "LicenseNo": LN, "Status": "Busy"}
								DriverAPI.UpdateStatus(strconv.Itoa(DriverID), jsonDataS)
								// Initiative End Trip
								fmt.Print("Distination Reached?\n", "Enter 1 to End Trip: ")
								fmt.Scanln(&end)
								if end == 1 {
									// Map Trip Details into string:interface as TripID datatype is integer
									jsonData := map[string]interface{}{"TripID": intTrip, "TripEndTime": string(time.Now().Format("01-02-2006 15:04:05")), "Status": "Completed"}
									TripAPI.UpdateTrip(TripOption, jsonData, "end")
									// Map Driver Details into string:interface as DriverID is integers
									jsonDataS := map[string]interface{}{"DriverID": DriverID, "FirstName": FN, "LastName": LN,
										"MobileNo": MN, "EmailAddress": EA, "LicenseNo": LN, "Status": "Available"}
									DriverAPI.UpdateStatus(strconv.Itoa(DriverID), jsonDataS)
									break
								}
							}
						} else {
							fmt.Println("\nThere are no Pending Trips!")
							break
						}
					}
				} else if DriverOption == 4 {
					// Passenger Option 4 - Back
					break
				} else {
					// Driver Menu Invalid Option
					fmt.Println("\nInvalid Option!")
				}
			}
		}
	}
}
