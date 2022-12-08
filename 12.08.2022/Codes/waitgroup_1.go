// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup //wg->package scope, sync->package
// A WaitGroup waits for a collection of goroutines to finish.

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1) // The main goroutine calls Add to set the number of
	go foo()
	bar()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait() // goroutines to wait for. Then each of the goroutines
}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
	// runs and calls Done when finished. At the same time,
	// Wait can be used to block until all goroutines have finished.
}
func bar() {
	for i := 0; i <= 10; i++ {
		fmt.Println("bar: ", i)
	}
}
