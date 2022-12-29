package main

import "fmt"

func main() {
	a := 2
	b := 0
	
	Divide(a, b)
}

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("can't divide '%d' by zero", a)
    }
    return a / b, nil
}
