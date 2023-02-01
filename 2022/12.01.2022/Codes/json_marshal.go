// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Javokhir",
		Last:  "Nematov",
		Age:   20,
	}

	p2 := person{
		First: "Ravi",
		Last:  "Seyed-Mahmoud",
		Age:   35,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
