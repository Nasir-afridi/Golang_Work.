package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("is Adult.")
	} else if age >= 12 {
		fmt.Println("Person is teenager.")
	} else {
		fmt.Println("person is kid.")
	}

	// operators
	role := "admin"
	hasPermission := true

	if role == "admin" || hasPermission {
		fmt.Println("Yes")
	}

	// declaring a variable in if condition. we can use this in else if also.
	if name := "Nasir"; name == "Nasir" {
		fmt.Println("Correct")
	}

	// Go mein ternary operator nahe hai.
}
