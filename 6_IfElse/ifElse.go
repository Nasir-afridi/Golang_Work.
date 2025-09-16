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
}
