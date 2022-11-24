// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type person struct {
	first  string
	second string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.second)
}

func main() {
	sa := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	fmt.Println(sa)
	sa.speak()
}
