package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "123456789"

	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	// bs (hash) is ready
	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not Logged in")
	}

	log.Println("Logged in!")
}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while generating bcrypt hash from password: %w", err)
	}

	return bs, nil
}

func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("invalid Password: %w", err)
	}
	return nil
}

// Continued on "Exploring-HMAC" folder