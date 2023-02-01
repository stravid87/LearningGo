// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var name string

func main() {
	name := "Facebooks"
	switch name {
	case "Apple":
		fmt.Println("Stiv Jobs")
	case "Amazon":
		fmt.Println("Jeff Bezos")
	case "Facebook":
		fmt.Println("Mark Zuckerberg")
	default:
		fmt.Println("Nothing")
	}
}
