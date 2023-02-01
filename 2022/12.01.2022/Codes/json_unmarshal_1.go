// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json: "First"`
	Last  string `json: "Last"`
	Age   int    `json: "Age"`
}

func main() {
	s := `[{"First":"Javokhir","Last":"Nematov","Age":20},{"First":"Ravi","Last":"Seyed-Mahmoud","Age":35}]`
	bs := []byte(s)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for _, v := range people {
		fmt.Println(v.First, v.Last, v.Age)
	}
}
