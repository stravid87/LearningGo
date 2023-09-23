package main

import (
	"strconv"
	"syscall/js"
	"strings"
)

func addTwoNumbers(this js.Value, args []js.Value) interface{} {
	a := args[0].Int()
	b := args[1].Int()
	sum := a + b
	return js.ValueOf(sum)
}

func formatNumberWithCommas(this js.Value, args []js.Value) interface{} {
	n := args[0].Int()
	// Convert the number to a string
	numStr := strconv.Itoa(n)

	// Split the string into parts
	parts := []string{}

	// Start from the end of the string and process it in groups of three digits
	for len(numStr) > 3 {
		parts = append(parts, numStr[len(numStr)-3:])
		numStr = numStr[:len(numStr)-3]
	}

	// Add the remaining digits
	parts = append(parts, numStr)

	// Reverse the parts slice
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}

	// Join the parts with a space separator
	result := strings.Join(parts, " ")

	return result
}

func main() {
	c := make(chan struct{}, 0)
		js.Global().Set("addTwoNumbers", js.FuncOf(addTwoNumbers))
		js.Global().Set("formatNumberWithCommas", js.FuncOf(formatNumberWithCommas))
	<-c
}
