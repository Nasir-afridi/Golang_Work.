package main

import (
	"fmt"
	"maps"
)

func main() {
	//creating map.
	m := make(map[string]string) // square bracket we define the key type and other one is value type.

	//setting an element.
	m["name"] = "golang"
	m["area"] = "Backend"

	// get an element.
	fmt.Println(m["name"])
	fmt.Println(m["area"])

	// if key does not exists in the map then it return the zero value.
	fmt.Println(m["phone"])

	// using the int string key with int
	a := make(map[string]int)
	a["age"] = 30
	a["number"] = 9999999
	fmt.Println(a["age"])
	fmt.Println(a["number"])

	// get the length
	fmt.Println(len(a))

	// delete the element from the map.
	delete(a, "number")
	fmt.Println(a)

	// clearing the map.
	clear(a)
	fmt.Println(a)

	// Map without using the make function
	x := map[string]int{"price": 1000, "age": 23}
	fmt.Println(x["price"])

	// checking two maps are equal or not.
	x1 := map[string]int{"price": 1000, "age": 23}
	x2 := map[string]int{"price": 1000, "age": 23}
	fmt.Println(maps.Equal(x1, x2))

}
