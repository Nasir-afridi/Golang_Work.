package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch.

	i := 5
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	}

	// Multiple Condition switch.
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's working days.")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("its integer")
		case string:
			fmt.Println("It's string")
		case bool:
			fmt.Println("bool")
		default:
			fmt.Println("other")
		}

	}

	whoAmI(6.444)

}
