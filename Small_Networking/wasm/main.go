package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
)

func main() {
	c := make(chan struct{})
	js.Global().Set("getEncryptedData", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		promiseConstructor := js.Global().Get("Promise")
		promise := promiseConstructor.New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]
			go func() {
				text, err := grabText()
				if err != nil {
					fmt.Println(err)
					reject.Invoke(js.ValueOf(err.Error()))
					return
				}

				key := []byte("this-is-a-32-byte-long-key-!!-rj")
				ciphertext, _ := hex.DecodeString(text)
				decryptedText, err := decrypt(ciphertext, key)
				if err != nil {
					fmt.Println(err)
					reject.Invoke(js.ValueOf(err.Error()))
					return
				}

				resolve.Invoke(string(decryptedText))
			}()
			return nil
		}))

		return promise
	}))
	<-c
}

func grabText() (string, error) {
	url := "http://localhost:9091"
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to process http request: %w", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read http body: %w", err)
	}
	return string(bodyBytes), nil
}

func decrypt(ciphertext []byte, key []byte) (string, error) {
    c, err := aes.NewCipher(key)
    if err != nil {
        return "error: ", err
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        return "error: ", err
    }

    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return "error: ", errors.New("ciphertext too short")
    }

    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    byteText, err := gcm.Open(nil, nonce, ciphertext, nil)


	return string(byteText), err

}