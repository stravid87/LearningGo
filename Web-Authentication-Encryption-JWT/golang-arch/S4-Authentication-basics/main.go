package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("ravi:123456789")))
}

// Continued on Password-storage folder