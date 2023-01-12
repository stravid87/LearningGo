package main 

import "fmt"

func main() {
	fmt.Println("Hello")
	f := foo()
	fmt.Println(f())
}

func foo() func() int {
	return func() int{
		return 64
	}
}