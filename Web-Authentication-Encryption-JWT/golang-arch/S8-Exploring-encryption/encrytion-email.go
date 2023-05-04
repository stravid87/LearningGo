package main

import (
	"bytes"
	"crypto/aes" // AES - Advanced Encryption Standard
	"crypto/cipher"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	msg := "There are few couple potentially more A than the duo of Brad Pitt and Angelina Jolie who met on the set of Mr & Mrs Smith."
	password := "ravi's_birthday"
	
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	//-->"bcrypt.GenerateFromPassword": This is a function from the bcrypt package that takes a plain-text password and a cost
	//   as input parameters and returns a hashed version of the password
	//-->"[]byte(password)": This is converting the plain-text password (assumed to be a string) into a byte slice, which is the
	//	 expected input format for the GenerateFromPassword function
	//-->"bcrypt.Mincost": This is constant from bcrypt package that represents the minimum allowed cost for hashing a password.
	//   The cost determines the complexity of the hashing algorithm, with higher costs resulting in slower but more secure hashes.
	if err != nil {
		log.Fatalln("couldn't bcrypt password", err)
	}

	bs := hashedPass[:16]

	rslt, err := enDecode(bs, msg)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("before base64", string(rslt))

	rslt2, err := enDecode(bs, string(rslt))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rslt2))
}

func enDecode(key []byte, input string) ([]byte, error) {
	// 1.Creates a new AES cipher with the given key. If there is an error creating the cipher, it returns nil and error message
	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("couldn't newCipher: %w", err)
	}

	// 2.Initializes an empty byte slice with the length of AES block size(16 bytes) to be used as the inizialization vector
	iv := make([]byte, aes.BlockSize)
	// 3.Creates a new CTR mode stream cipher using the AES cipher b and the initialization vector iv.
	s := cipher.NewCTR(b, iv)
	
	// 4.Inizialize a new buffer to store the encrypted or decrypted data.
	buff := &bytes.Buffer{}
	// 5.Creates a new cipher.StreamWriter that uses CTR mode stream cipher s and writes the output to the buffer buff
	sw := cipher.StreamWriter {
		S: s,
		W: buff,
	}
	// 6.Encrypts or decrypts the input string by writing it to the cipher.Stream.Writer
	_, err = sw.Write([]byte(input))
	if err != nil {
		return nil, fmt.Errorf("couldn't sw.Write to streamwriter %w", err)
	}
	// 7.Returns the encrypted or decrypted data as a byte slice
	return buff.Bytes(), nil
}