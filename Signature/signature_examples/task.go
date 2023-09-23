package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto"
	"encoding/base64"
	"fmt"
)

func sign(privateKey *rsa.PrivateKey, message string) (string, error) {
	// Hash the message using SHA256
	hash := sha256.Sum256([]byte(message))

	// Sign the hash using the private key
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign the hash: %v", err)
	}

	// Convert signature bytes to base64 string
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	return signatureBase64, nil
}

func verify(publicKey *rsa.PublicKey, message string, signatureBase64 string) error {
	// Hash the message using SHA256
	hash := sha256.Sum256([]byte(message))

	// Convert base64 signature string back into bytes
	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return fmt.Errorf("failed to decode base64 signature: %v", err)
	}

	// Verify the signature
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return fmt.Errorf("signature verification failed: %v", err)
	}

	return nil
}

func main() {
	// Declare a message
	message := "Hello, world!"

	// Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Failed to generate RSA key pair:", err)
		return
	}

	// Sign the message using the private key
	signature, err := sign(privateKey, message)
	if err != nil {
		fmt.Println("Failed to sign the message:", err)
		return
	}

	// Verify the signature using the public key
	err = verify(&privateKey.PublicKey, message, signature)
	if err != nil {
		fmt.Println("Signature verification failed:", err)
		return
	}

	fmt.Println("Signature verification successful!")
}
