package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", mySumchanged(2, 3))
	fmt.Println("4 + 7 =", mySumchanged(4, 7))
	fmt.Println("6 + 3 =", mySumchanged(6, 3))
}

func mySumchanged(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}