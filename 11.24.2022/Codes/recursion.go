// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	f := factorial(4)
	fmt.Println(f)

	sum := forFactorial(4)
	fmt.Println(sum)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func forFactorial(n int) int {
	s := 1
	for ; n > 0; n-- {
		s *= n
	}
	return s
}
