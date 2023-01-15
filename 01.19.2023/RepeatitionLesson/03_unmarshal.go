package main

import (
	"encoding/json"
	"fmt"
)
type person struct{
	First string `json:"First"`
	Age int `json:"Age"`
}
func main() {
	j := `[{"First":"James Bond","Age":34},{"First":"Miss Moneypeny","Age":24}]`
	fmt.Println(j)

	var people []person

	err := json.Unmarshal([]byte(j), &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	fmt.Println("\t", "First name", "Age")
	for _, v := range people{
		fmt.Println("\t", v.First, v.Age)
	}
}