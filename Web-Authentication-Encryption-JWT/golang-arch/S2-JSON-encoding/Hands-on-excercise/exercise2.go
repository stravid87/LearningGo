package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Javokhir",
	}
	p2 := person{
		First: "Ravi",
	}

	people := []person{p1, p2}
	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Println("ENCODING ERROR", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	people := []person{}
	err := json.NewDecoder(r.Body).Decode(&people)
	if err != nil {
		log.Println("Decoding error", err)
	}

	fmt.Println("Person", people)
}