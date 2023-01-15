package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First string
	Last string
	Age int
	Arr []string
}

func main() {
	u1 := person{
		First: "James",
		Last: "Bond",
		Age: 45,
		Arr: []string{
			"This is text",
			"This is another text",
			"I am a text",
		},
	}
	u2 := person{
		First: "Miss",
		Last: "MoneyPenny",
		Age: 27,
		Arr: []string{
			"This is text",
			"This is another text",
			"I am a text",
		},
	}

	users := []person{u1, u2}
	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {
		fmt.Println(err)
	}
}