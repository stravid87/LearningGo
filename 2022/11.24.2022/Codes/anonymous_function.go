// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// Anonymous function
func main() {
	func() {
		fmt.Println("I am an Anonymous Function")
	}()

	func(x int) {
		fmt.Println("X is:", x)
	}(42)
}
