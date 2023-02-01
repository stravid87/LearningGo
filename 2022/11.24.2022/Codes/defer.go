// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	defer first()
	second()
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}
