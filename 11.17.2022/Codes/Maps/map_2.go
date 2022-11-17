// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	ab := map[string]int{
		"Javokhir": 20,
		"Ravi":     35,
	}

	delete(ab, "Javokhir")

	fmt.Println(ab)
}
