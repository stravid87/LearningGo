// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "Javokhir",
		last:  "Nematov",
	}

	p2 := person{
		first: "Ravi",
		last:  "Sayyid",
	}
	fmt.Println(p1, p2)
}
