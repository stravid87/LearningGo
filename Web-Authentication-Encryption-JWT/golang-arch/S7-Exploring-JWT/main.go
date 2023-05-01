package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true){
		return fmt.Errorf("Token has expired")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("Invalid Session Id")
	}

	return nil
}

var key = []byte{}

func main() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}
	pass := "123456789"

	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

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

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)

	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("error in signMessage while hashing message: %w", err)
	}

	signature := h.Sum(nil)
	return signature, nil
}

func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("error in checking while getting signature of message: %w", err)
	}

	same := hmac.Equal(newSig, sig)
	return same, nil
}

func createToken(c*UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodES512, c)
	signedToken, err := t.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("Error in createToekn when signing token: %w", err)
	}
	return signedToken, nil
}