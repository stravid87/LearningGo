// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// Anonymous function
func main() {
	a := foo()
	fmt.Println(a)

	b := bar()
	fmt.Println(b()())
}

func foo() string {
	s := "hello world"
	return s
}

func bar() func() int {
	return func() int {
		return 473
	}
}
