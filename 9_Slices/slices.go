package main

import "fmt"

func main() {
	// slices : ye fixed size ki nahe hoteen ye dynamic hoteen hain.
	// uninitialized slice is nill.

	var nums []int // ismein array size nahe mention krty ham.
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var Nums = make([]int, 2) // initial size of slice is 2 its not fixed. iski capacity or length automatically resize hojate hai.
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
	fmt.Println(Name)
}
