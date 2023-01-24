package main

import (
	"fmt"
)

func main(){
	c := make(chan int)

	go func(){
		for i := 1; i < 100; i++{
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}