package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter the number")
	fmt.Scanln(&num)

	switch num {
	case 1:
		fmt.Println("saturday")
	case 2:
		fmt.Println("sunday")
	case 3:
		fmt.Println("Monday")
	case 4:
		fmt.Println("Tuesday")
	case 5:
		fmt.Println("Wednesday")
	case 6:
		fmt.Println("Thursday")
	case 7:
		fmt.Println("Friday")
	default:
		fmt.Println("Enter the correct number")
	}
}
