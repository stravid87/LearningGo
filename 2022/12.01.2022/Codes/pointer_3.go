// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	a := 34

	b := &a
	fmt.Println(*b)
	fmt.Println(*&a)

	*b = 45
	fmt.Println(a)
}
