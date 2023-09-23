package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

type Transaction struct {
	Sender    string
	Receiver  string
	Amount    float64
	Signature []byte
}

func Sign(privateKey *rsa.PrivateKey, data []byte) ([]byte, error) {
	hash := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func Verify(publicKey *rsa.PublicKey, data []byte, signature []byte) bool {
	hash := sha256.Sum256(data)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return false
	}
	return true
}

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	publicKey := &privateKey.PublicKey

	transaction := Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   1.5,
	}

	data := []byte(transaction.Sender + transaction.Receiver + fmt.Sprintf("%.2f", transaction.Amount))

	signature, err := Sign(privateKey, data)
	if err != nil {
		fmt.Println("Error signing transaction:", err)
		return
	}

	transaction.Signature = signature

	valid := Verify(publicKey, data, transaction.Signature)
	fmt.Println("Transaction Signature Valid:", valid)
}
