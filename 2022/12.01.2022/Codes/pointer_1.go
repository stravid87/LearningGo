// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	a := 34
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(*b)
}
