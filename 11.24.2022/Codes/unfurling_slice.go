// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	fmt.Println(x)
}
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for v := range x {
		sum += v
	}
	return sum
}
