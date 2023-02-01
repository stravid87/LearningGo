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

	m := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}

	for _, v := range m{
		fmt.Println(v.first)
		fmt.Println(v.last)

		for i, val := range v.favourite{
			fmt.Println(i, val)
		}
	}
}