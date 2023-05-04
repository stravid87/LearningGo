package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	msg := "There are few couple potentially more A than the duo of Brad Pitt and Angelina Jolie who met on the set of Mr & Mrs Smith."
	encoded := encode(msg)
	fmt.Println("ENCODED", encoded)


	s, err := decode(encoded)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("DECODED", s)
}

func encode(msg string) string {
	return base64.URLEncoding.EncodeToString([]byte(msg))
}

func decode(encoded string) (string, error) {
	s, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return "", fmt.Errorf("couldn't decode string %w", err)
	}
	return string(s), nil
}