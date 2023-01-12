package main

import "fmt"

func main(){
	f := func(){
		for i := 1; i < 4; i++{
			fmt.Println(i)
		}
	}

	f()
}