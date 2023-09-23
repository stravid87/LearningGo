package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	// "crypto/x509"
	// "encoding/pem"
	"fmt"
	"log"
)

func main() {
	// Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	// Create a dummy message
	message := []byte("Hello, world!")

	// Sign the message
	signature, err := sign(message, privateKey)
	if err != nil {
		log.Fatalf("Failed to sign message: %v", err)
	}

	// Verify the signature
	valid := verify(message, signature, &privateKey.PublicKey)
	if valid {
		fmt.Println("Signature is valid")
	} else {
		fmt.Println("Signature is invalid")
	}
}

func sign(message []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	// Hash the message
	hashed := sha256.Sum256(message)

	// Sign the hashed message
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, fmt.Errorf("Failed to sign message: %v", err)
	}

	return signature, nil
}

func verify(message, signature []byte, publicKey *rsa.PublicKey) bool {
	// Hash the message
	hashed := sha256.Sum256(message)

	// Verify the signature
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return false
	}

	return true
}
