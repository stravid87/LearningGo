// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type person struct {
	first string
	last  string
}
type AllStruct struct {
	person
}

func main() {
	sa := AllStruct{
		person: person{
			first: "javokhir",
			last:  "Nematov",
		},
	}
	fmt.Println(sa.person.last)
}
