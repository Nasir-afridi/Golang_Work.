package main

import "fmt"

func main() {
	var age int

	fmt.Println("Enter the age of the user : ")
	fmt.Scanln(&age)

	if age < 18 {
		fmt.Println("Minor")
	} else if age >= 18 && age < 60 {
		fmt.Println("Adult")
	} else if age >= 60 {
		fmt.Println("Senior citizen")
	}
}
