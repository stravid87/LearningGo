// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutine:", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() //Locked Goroutines
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock() // Unlock Goroutines
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutine:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)
}
