// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	fmt.Println( runtime.NumGoroutine())
	teams := []string{"Ravi", "Arnon", "Hermann", "Javokhir"}

	for _, team := range teams {
		go attack(team)
	}
}
func attack(team string) {
	fmt.Println("Start team ", team)
	time.Sleep(time.Second)
}
