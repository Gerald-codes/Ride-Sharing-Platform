package main

import (
	"fmt"
	DriverMain "importMods/Driver_MS"
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

		} else if UserOption == 2 {
			fmt.Println("----------Enter Driver Details----------")
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
			jsonData := map[string]interface{}{"IC No": NRIC, "First Name": FN, "Last Name": LN,
				"Mobile No": MN, "Email Address": EA, "License No": LiN, "Status": "Available"}
			DriverMain.AddDriver("1", jsonData)
			DriverMain.GetDriver("1")
			DriverMain.AddDriverToDB(jsonData)
		}
	} else if MenuOption == 2 {
		fmt.Println("-------------Edit Account--------------", User[0]+User[1])
		fmt.Println("Option: ")
		fmt.Scanln(&UserOption)
		if UserOption == 1 {

		} else if UserOption == 2 {
			fmt.Println("----------Enter Driver Details----------")
			var NRIC string
			fmt.Println("NRIC Number: ")
			fmt.Scanln(&NRIC)
			jsonData := map[string]interface{}{"IC No": "T01212193L", "First Name": "Gerald", "Last Name": "Tan",
				"Mobile No": 91234567, "Email Address": "Testing@test.com", "License No": "LKS123L", "Status": "Free"}
			DriverMain.UpdateDriver("1", jsonData)
			DriverMain.GetDriver("1")
		}
	}
}
