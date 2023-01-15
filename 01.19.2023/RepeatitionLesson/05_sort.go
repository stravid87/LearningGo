package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{3, 5, 7, 1, 7, 8, 6, 23, 43, 23, 12, 15}
	xs := []string{"v", "d", "e", "t", "a", "e", "b", "g"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}