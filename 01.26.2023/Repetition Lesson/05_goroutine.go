package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	incrementor := 0
	gs := 100
	wg.Add(100)

	for i := 0; i < gs; i++{
		go func(){
			m.Lock()
				v := incrementor
				v ++
				incrementor = v
			m.Unlock()
			wg.Done()
		}()
	} 
	wg.Wait()
	fmt.Println("end", incrementor)
}