package main

import (
	"fmt"
	"https://github.com/globe-and-citizen/LearningGo/tree/main/01.05.2023/Codes"
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