package UIMain

import (
	"fmt"
	"/Users/gerald/Desktop/SEM 3.2/ETI/Assignment/Ride-Sharing-Platform/src/Driver_MS"
)

func UIMain() {
	var MenuOption, CreateAccMenuOption float32
	var Menu [3]string
	Menu[0] = "\n[1] Create Account" // Assign a value to the first element
	Menu[1] = "\n[2]"                // Assign a value to the second element
	Menu[2] = "\n[3]"                // Assign a value to the third element
	var CreateAccMenu [3]string
	CreateAccMenu[0] = "\n[1] Passenger" // Assign a value to the first element
	CreateAccMenu[1] = "\n[2] Driver"    // Assign a value to the second element

	fmt.Println("---------Welcome to Ride Sharing---------", Menu[0]+Menu[1]+Menu[2])
	fmt.Println("Option: ")
	fmt.Scanln(&MenuOption)
	if MenuOption == 1 {
		// s := "echo helloasdasd"
		// args := strings.Split(s, "")
		// cmd := exec.Command(args[0], args[1:]...)
		// b, err := cmd.CombinedOutput()
		// if err != nil {
		// 	fmt.Printf("asdasd %v", err)
		// }
		// fmt.Printf("%s", b)
		fmt.Println("-------------Create Account--------------", CreateAccMenu[0]+CreateAccMenu[1])
		fmt.Println("Option: ")
		fmt.Scanln(&CreateAccMenuOption)
		if CreateAccMenuOption == 1 {
			DriverMain.hello()
		}
	}

}
