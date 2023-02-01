package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 34

	fmt.Println(<-c)
}