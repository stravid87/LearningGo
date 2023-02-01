package main

import "fmt"

func main() {
	for i := 10; i <= 100; i ++ {
		fmt.Printf("Original is %v, Dividev value is %v\n", i, i%4)
	}
}