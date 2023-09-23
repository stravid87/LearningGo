package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
)

type SimplePost struct {
	Signature string
	PublicKey *rsa.PublicKey
	Message []byte
}

func SignString(this js.Value, args []js.Value) interface{} {
	arg0BS := []byte(args[0].String())
	resolve_reject_internals := func(this js.Value, args []js.Value) interface{} {
		resolve := args[0]
		reject := args[1]
		go func(input []byte) {
			privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
			if err != nil {
				fmt.Errorf("Error on POST to: %s", err.Error())
				reject.Invoke(js.ValueOf("Failure on Post"))
			}

			hash := sha256.Sum256([]byte(input))
			signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
			if err != nil {
				reject.Invoke(js.ValueOf(fmt.Errorf("Error hashing: \"%s\"", err.Error())))
			}

			signatureBase64 := base64.StdEncoding.EncodeToString(signature)
			var url = "http://localhost:9090/post-signature"
			simplePost := SimplePost{
				Signature:  signatureBase64,
				PublicKey: &privateKey.PublicKey,
				Message: input,
			}

			simplePost_bs, err := json.Marshal(simplePost)
			if err != nil {
				fmt.Errorf("Error on POST to %s: %s", url, err.Error())
				reject.Invoke(js.ValueOf("Failure on Post"))
			}

			resp, err := http.Post(url, "Content-Type:application/json", bytes.NewReader(simplePost_bs))
			if err != nil {
				fmt.Errorf("Error on POST to %s: %s", url, err.Error())
				reject.Invoke(js.ValueOf("Failure on Post"))
			}

			response_BS, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				reject.Invoke(js.ValueOf(fmt.Errorf("Error reading response body: ", err.Error())))
			}

			resolve.Invoke(js.ValueOf(fmt.Sprintf(string(response_BS))))
		}(arg0BS)
		return nil
	}
	promiseConstructor := js.Global().Get("Promise")
	promise := promiseConstructor.New(js.FuncOf(resolve_reject_internals))
	return promise
}