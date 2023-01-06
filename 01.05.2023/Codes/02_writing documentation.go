package main

import (
	"fmt"
	"github link"
)

type dogs struct{
	name string
	age int
}

func main() {
	dog := dogs{
		name: "Pet",
		age: 2,
	}
	fmt.Println(dog)
}