// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	for i := 65; i <= 99; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("%#U\n", i)
		}
	}
}
