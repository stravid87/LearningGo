package main

import (
	"fmt"
	"encoding/json"
	"bytes"
)

type user struct {
    Name string
    Age int
}

func main() {
    buf := new(bytes.Buffer)
    u := user{
        Name: "bob",
        Age: 20,
    }
    err := json.NewEncoder(buf).Encode(u)
    if err != nil {
        panic(err)
    }
    fmt.Print(buf.String())
}