package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type person struct {
	First string
	Last string
	Sayings []string
}

func main() {
	p1 := person{
		First: "James",
		Last: "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("Json did not marshal - here's the error:", err)
	}
	fmt.Println(string(bs))
}