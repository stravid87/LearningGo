package main

import (
	_ "crypto/sha512"
	"log"
	"time"

	"fmt"
	"syscall/js"

	"github.com/dgrijalva/jwt-go"
)

func main() {
    done := make(chan struct{}, 0)
    js.Global().Set("authorizationJWT", js.FuncOf(auth))
	<-done
}

type myClaims struct {
	jwt.StandardClaims
	Email string
}

const myKey = "I love thursday when it rains"

func auth(this js.Value, args []js.Value) interface{} {
    emailValue := args[0]
    email := emailValue.String()

	log.Println("email")

	ss, err := getJWT(email)
	if err != nil {
        return fmt.Sprintf("Error: couldn't getJWT: %v", err)
    }

    return ss
}

func getJWT(msg string) (string, error) {
    claims := myClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
		Email: msg,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	ss, err := token.SignedString([]byte(myKey))
	if err != nil {
		return "", fmt.Errorf("couldn't SignedString %w", err)
	}
	return ss, nil
}