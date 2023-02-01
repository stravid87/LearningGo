// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	x := 34
	foo(x)
	fmt.Println(x)
}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}
