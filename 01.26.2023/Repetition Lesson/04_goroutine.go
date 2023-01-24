package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	var wg sync.WaitGroup

	incrementor := 0
	gs := 100
	wg.Add(100)

	for i := 0; i < gs; i++{
		go func(){
			v := incrementor
			runtime.Gosched()
			v ++
			incrementor = v
			wg.Done()
		}()
	} 
	wg.Wait()
	fmt.Println("end", incrementor)
}