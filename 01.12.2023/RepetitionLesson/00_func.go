package main

import "fmt"

func main() {
	n := foo()
	x, s := bar()

	fmt.Println(n, x, s)
}

func foo() int {
	return 20
}

func bar() (int, string) {
	return 2002, "Go is the best Language!"
}