package main

import "fmt"

func main() {
	// for -> sirf for loop he hota hai golang mein. for ki madad sy while or do while banatay hain.

	// while loop.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// infinite loop with break example.
	for {
		fmt.Println("hello")
		break // exit after one iteration to avoid unreachable code
	}

	// Classic for loop yane for loop,
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}
}
