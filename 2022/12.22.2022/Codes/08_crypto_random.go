package main

import "fmt"
import "crypto/rand"

func main(){
	randomCrypto, _ := rand.Prime(rand.Reader, 128)
	fmt.Println(randomCrypto)
}