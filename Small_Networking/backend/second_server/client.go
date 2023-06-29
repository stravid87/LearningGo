package main

import (
	// "bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"time"
	"fmt"
	"io"
	"log"
	"net/http"
	"proto"

	"github.com/rs/cors"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	// grpc.Dial is a function that creates a client connection to the given target
	// "localhost:9090" is the target which the function is creating a connection to.
	// grpc.WithInsecure() is an option to create the connection without encryption, which
	// is not recommended for production code, but it's fine in development or testing stage when security is not the critical consideration
	if err != nil {
		log.Println(err)
	}

	client := proto.NewChatServiceClient(connection)

	message := proto.Message{
		Body: "Thank you for long Lorem!",
	}

	resp, err := client.SendLorem(context.Background(), &message)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp.Body) // Random text

	text := []byte(resp.Body)
	key := []byte("this-is-a-32-byte-long-key-!!-rj")
	
	cipherText, err := encrypt(text, key)
	if err != nil {
		fmt.Println(err)
	}
	
	startServer(cipherText)
}

func encrypt(plaintext []byte, key []byte) ([]byte, error) {

	// Validate key length
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, fmt.Errorf("crypto/aes: invalid key size %d, want: 16, 24 or 32", len(key))
	}

	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, nil
}

func startServer(cipherText []byte) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%x", cipherText)
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Language", "Content-Type"},
	})

	server:= &http.Server{
		Addr: ":9091",
		ReadTimeout: 5 * time.Minute,
		WriteTimeout: 10 * time.Second,
		Handler: c.Handler(mux),
	}

	log.Fatal(server.ListenAndServe())
}