package main

import (
	"fmt"
)

func main() {
	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	// The fmt.Scan() function in Go language scans the input texts which is given in the standard input,
	// reads from there and stores the successive space-separated values into successive arguments. 
	// Moreover,this function is defined under the fmt package. Here, you need to import the “fmt” package in order to use these functions.
	// In addition, We can save data to real database with Scan()
	_, err := fmt.Scan(&answer1)
	// If you want to see how error works, fill integer for input
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Fav food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(answer1, answer2, answer3)
}