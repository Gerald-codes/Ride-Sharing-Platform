package main

import (
	"fmt"
	DriverAPI "importMods/UI/API/Driver"
	PassengerAPI "importMods/UI/API/Passenger"
)

func main() {
	var MenuOption, UserOption float32
	var Menu [3]string
	Menu[0] = "\n[1] Create Account" // Assign a value to the first element
	Menu[1] = "\n[2] Edit Account"   // Assign a value to the second element
	Menu[2] = "\n[3] View All"       // Assign a value to the third element
	var User [3]string
	User[0] = "\n[1] Passenger" // Assign a value to the first element
	User[1] = "\n[2] Driver"    // Assign a value to the second element

	fmt.Println("---------Welcome to Ride Sharing---------", Menu[0]+Menu[1]+Menu[2])
	fmt.Println("Option: ")
	fmt.Scanln(&MenuOption)
	if MenuOption == 1 {
		fmt.Println("-------------Create Account--------------", User[0]+User[1])
		fmt.Println("Option: ")
		fmt.Scanln(&UserOption)
		if UserOption == 1 {
			jsonData := map[string]interface{}{"First Name": "Troll", "Last Name": "Ta1n",
				"Mobile No": 91234567, "Email Address": "Testing@test.com"}
			PassengerAPI.AddPassenger("91234567", jsonData)
		} else if UserOption == 2 {
			fmt.Println("----------Enter Driver Details----------")
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
			// DriverAPI.GetDriver("1")
		}
	} else if MenuOption == 2 {
		fmt.Println("-------------Edit Account--------------", User[0]+User[1])
		fmt.Println("Option: ")
		fmt.Scanln(&UserOption)
		if UserOption == 1 {
			jsonData := map[string]interface{}{"First Name": "GERALD", "Last Name": "Ta1n",
				"Mobile No": 91234567, "Email Address": "Testing@test.com"}
			PassengerAPI.UpdatePassenger("91234567", jsonData)
		} else if UserOption == 2 {
			fmt.Println("----------Enter Driver Details----------")
			// var NRIC string
			// fmt.Println("NRIC Number: ")
			// fmt.Scanln(&NRIC)
			jsonData := map[string]interface{}{"NRIC No": "T0112123R", "First Name": "Gerald", "Last Name": "Tan",
				"Mobile No": 91234567, "Email Address": "Testing@test.com", "License No": "LKS123L", "Status": "Free"}
			DriverAPI.UpdateDriver("T0112123R", jsonData)
		}
	}
}
