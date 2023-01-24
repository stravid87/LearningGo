package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(2)
	go func(){
		fmt.Println("Hello from one!")
		wg.Done()
	}()

	go func(){
		fmt.Println("Hello from two!")
		wg.Done()
	}()

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("exited")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}