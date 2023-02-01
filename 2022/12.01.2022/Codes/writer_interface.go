// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello Javokhir")
	fmt.Fprintln(os.Stdout, "Hello Javokhir")
	io.WriteString(os.Stdout, "Hello Javokhir")
}
