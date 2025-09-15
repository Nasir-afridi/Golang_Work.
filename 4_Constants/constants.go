package main

import "fmt"

func main() {
	// we cannot declare a variable with shorthand syntax out of the function.
	const name = "nasir" // abb ham iski value change nahe krskty.
	fmt.Println(name)

	//grouped multiple constant together.
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(host)
}
