package main

import (
	"fmt"
	"os"
)

func main() {
    f, err := os.OpenFile("123.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    n, err := f.Write([]byte("I am being written to txt file!"))
    if err != nil {
        panic(err)
    }
    fmt.Println("wrote %d bytes", n)
}
