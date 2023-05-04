package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	msg := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. "
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