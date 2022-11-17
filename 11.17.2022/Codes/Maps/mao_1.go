package main

import "fmt"

func main() {
	m := map[string]int{
		"Javokhir": 20,
		"Ravi":     35,
	}
	fmt.Println(m)
	fmt.Println(m["Javokhir"])
	fmt.Println(m["Ravi"])
	fmt.Println(m["Arnon"])

	m["Herman"] = 25

	for k, v := range m {
		fmt.Println(k, v)
	}
}
