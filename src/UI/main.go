package main

import (
	"fmt"
	DriverMain "importMods/Driver_MS"
)

func main() {
	var MenuOption, CreateAccMenuOption float32
	var Menu [3]string
	Menu[0] = "\n[1] Create Account" // Assign a value to the first element
	Menu[1] = "\n[2] Edit Account"   // Assign a value to the second element
	Menu[2] = "\n[3] View All"       // Assign a value to the third element
	var CreateAccMenu [3]string
	CreateAccMenu[0] = "\n[1] Passenger" // Assign a value to the first element
	CreateAccMenu[1] = "\n[2] Driver"    // Assign a value to the second element

	fmt.Println("---------Welcome to Ride Sharing---------", Menu[0]+Menu[1]+Menu[2])
	fmt.Println("Option: ")
	fmt.Scanln(&MenuOption)
	if MenuOption == 1 {
		fmt.Println("-------------Create Account--------------", CreateAccMenu[0]+CreateAccMenu[1])
		fmt.Println("Option: ")
		fmt.Scanln(&CreateAccMenuOption)
		if CreateAccMenuOption == 1 {
			DriverMain.Hello()
		} else if CreateAccMenuOption == 2 {
			DriverMain.Hello()
			jsonData := map[string]string{"ID": "t1212213o", "First Name": "Gerald", "Last Name": "Tan",
				"Mobile No": "91234567", "Email Address": "Testing@test.com", "License No": "LKS123L", "Status": "Available"}
			DriverMain.AddDriver("1", jsonData)
			DriverMain.GetDriver("1")
			DriverMain.AddDriverToDB()
		}
	}
}
