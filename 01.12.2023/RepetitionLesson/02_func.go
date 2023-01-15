package main 

import (
	"fmt"
)

func main() {
	defer foo()
	fmt.Println("Hello, Playground")
}

func  foo() {
	defer func(){
		fmt.Println("Too defer")
	}()
	fmt.Println("foo ran")
}