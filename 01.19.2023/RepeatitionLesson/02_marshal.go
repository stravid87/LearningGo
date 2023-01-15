package main

import (
	"encoding/json"
	"fmt"
)

type user struct{
	First string
	Age int
}

func main() {
	u1 := user{
		First: "James Bond",
		Age: 34,
	}
	u2 := user{
		First: "Miss Moneypeny",
		Age: 24,
	}

	users := []user{u1, u2}
	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}