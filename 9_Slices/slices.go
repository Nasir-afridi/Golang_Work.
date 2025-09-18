package main

import (
	"fmt"
	"slices"
)

func main() {
	// slices : ye fixed size ki nahe hoteen ye dynamic hoteen hain.
	// uninitialized slice is nill.

	var nums []int // ismein array size nahe mention krty ham.
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var Nums = make([]int, 2) // initial length of slice is 2 its not fixed. iski capacity or length automatically resize hojate hai.
	fmt.Println(cap(Nums))    // capacity -> maximum numbers of elements can fit.
	Nums = append(Nums, 1)    // adding element in a slice.
	fmt.Println(Nums)

	// agrr slice ki capacity 5 hai or hamny 6 add krdiye to wo capacity ko automatic barha laita hai yane wo 5 sy 10 hojyga.
	var Num = make([]int, 0, 5) // 0 is the length means how many elements are visible and 5 is the total capacity of elements.
	Num = append(Num, 2, 3, 4, 5, 6, 5, 3, 5, 6, 3, 11)
	fmt.Println(Num)
	fmt.Println(cap(Num))

	// initializing the slice.
	Name := []string{"Nasir", "Ali"}
	Name[0] = "Khan" // changing the value
	fmt.Println(Name)

	// ---------------------- Slice Difference ----------------------
	//
	// var s []int
	// - Creates a nil slice
	// - len = 0, cap = 0, pointer = nil
	// - Not ready for direct indexing (s[0] panic)
	// - Memory allocate hogi jab pehli dafa append use karoge
	// - Common for simple cases
	//
	// s := make([]int, length, capacity)
	// - Creates a ready-to-use slice
	// - len = user defined, cap = user defined
	// - Pointer already allocated (not nil)
	// - Safe for indexing (agar index < len)
	// - Useful for performance & when size known
	//
	// --------------------------------------------------------------

	// copy one slice into another.
	var value = make([]int, 0, 5)
	value = append(value, 2)
	var value1 = make([]int, len(value))
	copy(value1, value)
	fmt.Println(value, value1)

	// slice operator
	// 0 index yane from sy 2 yane to tk k beech mein jtny element hoty hain onko print krty hain yane 1, 3.
	var Val = []int{1, 3, 4}
	fmt.Println(Val[0:2])
	fmt.Println(Val[:2]) // by default 0 sy start hoga yane 1, 3 print kryga
	fmt.Println(Val[0:]) // end tkk jtny elements hongy on saab ko print kryga

	// checking two slices are equal or not
	var Val1 = []int{1, 3, 4}
	fmt.Println(slices.Equal(Val, Val1))
}
