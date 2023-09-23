package main

import (
	"syscall/js"
)

func main() {
	wasmBlockingChan := make(chan struct{})
	js.Global().Set("signString", js.FuncOf(SignString))
	<-wasmBlockingChan
}