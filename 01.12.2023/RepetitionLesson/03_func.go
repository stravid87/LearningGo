package main

import "fmt"

type person struct{
	name string
	last string
	age int
}

func (p person) speak(){
	fmt.Println("I am", p.name, p.last, "and I am", p.age, "years old.")
}

func main() {
	p1 := person{
		name: "James",
		last: "Bond",
		age: 34,
	}

	p1.speak()
}