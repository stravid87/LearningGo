package main

import "fmt"

func main() {
	x := [5]int{12, 23, 45, 64, 33}
	for i, v := range x{
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}