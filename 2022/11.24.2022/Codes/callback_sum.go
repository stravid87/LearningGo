// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 8}
	a := sum(ii...)
	fmt.Println(a)
}

func sum(ii ...int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}
