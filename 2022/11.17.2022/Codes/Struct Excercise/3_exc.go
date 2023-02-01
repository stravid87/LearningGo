package main 

import "fmt"

func main() {
	s := struct{
		first string
		friends map[string]int
		favourite []string
	}{
		first: "James",
		friends: map[string]int{
			"MoneyPenny": 555,
			"Q": 444,
			"T": 777,
		},
		favourite: []string{
			"Martini", 
			"Water",
		},
	}

	fmt.Println(s.first)

	for i, v := range s.friends{
		fmt.Println("<<", i, v, ">>")
	}

	for i, v := range s.favourite{
		fmt.Println("<-", i, v,"->")
	}
}