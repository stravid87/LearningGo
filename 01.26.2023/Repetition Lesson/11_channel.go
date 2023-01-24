package main

import (
	"fmt"
	"os"
)

func main(){
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("program is exited")
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	os.Exit(10)
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func(){
		for i := 0; i < 100; i ++{
			c <- i
		}
		close(c)
	}()
	return c
}