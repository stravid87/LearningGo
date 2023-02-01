// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 8}
	a := sum(ii...)
	fmt.Println(a)

	b := even(sum, ii...)
	fmt.Println("even number", b)
}

func sum(ii ...int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

func even(f func(ii ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
