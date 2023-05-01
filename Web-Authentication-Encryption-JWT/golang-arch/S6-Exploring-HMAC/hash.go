package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	h := sha256.New()
	io.WriteString(h, "hello world")
	s := h.Sum(nil)
	fmt.Println(hex.EncodeToString(s))
}