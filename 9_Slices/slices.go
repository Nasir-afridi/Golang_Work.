package main

import "fmt"

func main() {
	// slices : ye fixed size ki nahe hoteen ye dynamic hoteen hain.
	// uninitialized slice is nill.

	var nums []int // ismein array size nahe mention krty ham.
	fmt.Println(nums == nil)
	fmt.Println(len(nums))
}
