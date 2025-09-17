package main

import "fmt"

// fized size k liye array use hogi.
// array mein ham paihly sy he define kraingy kk ismein itny elements aanay walay hain or wo aik he type kk hongy.

func main() {
	//declaring the array.
	var nums [4]int

	// checking the Array length
	fmt.Println(len(nums))

	// Adding element to the Array.
	nums[0] = 33
	fmt.Println(nums[0])

	// declaring and initializing the array.
	Names := [4]string{"Nasir", "Ali", "Khan", "Ahmad"}
	fmt.Println(Names)

	// 2d Arrays.
	// first one iski row hoge yane 2 row or second iski column yane 3 column
	Nums := [2][3]int{{3, 4, 3}, {5, 6, 3}}
	fmt.Println(Nums)
}
