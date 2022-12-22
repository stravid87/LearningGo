// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(4)
	go one()
	go two()
	go three()
	go four()

	wg.Wait()
	fmt.Println("Terminating Program")
}

func one() {
	wg.Done()
	fmt.Println("One")
}

func two() {
	wg.Done()
	fmt.Println("Two")
}

func three() {
	wg.Done()
	fmt.Println("Three")
}

func four() {
	wg.Done()
	fmt.Println("Four")
}
