// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

func main() {
	ab := []int{1, 4, 56, 3, 2, 6, 7, 4, 3, 7, -1}

	fmt.Println(ab)
	sort.Ints(ab)
	fmt.Println(ab)
}
