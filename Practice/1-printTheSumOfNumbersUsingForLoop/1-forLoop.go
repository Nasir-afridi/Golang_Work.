// print the sum of 1 to 10 using for loop.

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
