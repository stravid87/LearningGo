package main 

import "fmt"

type person struct{
	first string
	last string
	favourite []string
}

func main() {
	p1 := person{
		first: "James",
		last : "Bond",
		favourite: []string{
			"chocolate",
			"cake",
			"drink",
		},
	}

	p2 := person{
		first: "Miss",
		last: "Moneypenny",
		favourite: []string{
			"apple",
			"banana",
			"pineapple",
		},
	}

	fmt.Println(p1, p2)

	fmt.Println(p1.first)
	for i, v := range p1.favourite{
		fmt.Println("\t", i, v)
	}

	fmt.Println(p2.first)
	for i, v := range p2.favourite{
		fmt.Println("\t", i, v)
	}
}