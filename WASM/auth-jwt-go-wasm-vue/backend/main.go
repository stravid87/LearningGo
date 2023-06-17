package main

import (
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("./front"))
    http.Handle("/", fs)

    log.Println("Listening on http://localhost:5000/index.html")
    err := http.ListenAndServe(":5000", nil)
    if err != nil {
        log.Fatal(err)
    }
}