package main

import "fmt"

func main() {

	x := make([]string, 5, 5)
	x = []string{"one", "two", "three", "four"} //4 elements

	fmt.Println(cap(x)) // prints 4 instead of 5

	x = append(x, "five") // exceeds previous capacity of 4
	// thus allocates a new underlying array
	// with capacity 8

	fmt.Println(cap(x)) // Prints 8. Which should still be 5
	// if only make would have been used
}
