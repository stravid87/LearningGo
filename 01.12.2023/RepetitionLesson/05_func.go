package main

import "fmt"

func main() {
	func(){
		for i := 0; i <= 99; i++{
			fmt.Println(i)
		}
	}()
}