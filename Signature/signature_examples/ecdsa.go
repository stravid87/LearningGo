package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {
	// Generate ECDSA key pair
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Printf("Failed to generate private key: %v\n", err)
		return
	}

	// Create a dummy message
	message := []byte("Hello, world!")

	// Sign the message
	signature, err := sign(message, privateKey)
	if err != nil {
		fmt.Printf("Failed to sign message: %v\n", err)
		return
	}

	// Verify the signature
	valid := verify(message, signature, &privateKey.PublicKey)
	if valid {
		fmt.Println("Signature is valid")
	} else {
		fmt.Println("Signature is invalid")
	}
}

func sign(message []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	// Hash the message
	hash := sha256.Sum256(message)

	// Sign the hashed message
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return nil, fmt.Errorf("Failed to sign message: %v", err)
	}

	// Serialize the signature
	signature := append(r.Bytes(), s.Bytes()...)
	// The ... syntax after s.Bytes() is called a variadic parameter, 
	// which allows the function to accept a variable number of arguments.
	// In this case, it allows us to concatenate multiple byte slices together.

	return signature, nil
}

func verify(message, signature []byte, publicKey *ecdsa.PublicKey) bool {
	// Hash the message
	hash := sha256.Sum256(message)

	// Extract the r and s values from the signature
	rBytes := signature[:32]
	sBytes := signature[32:]
	r := new(big.Int).SetBytes(rBytes)
	s := new(big.Int).SetBytes(sBytes)

	// Verify the signature
	valid := ecdsa.Verify(publicKey, hash[:], r, s)

	return valid
}
