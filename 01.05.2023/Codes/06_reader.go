package main

import (
    "log"
    "os"
)

func main() {
    file, err := os.Open("123.txt")
    if err != nil {
        log.Fatalf("error opening sloth.txt: %v", err)
    }
    defer file.Close()

    // Make a byte slice that's big enough to store a few words of the message
    // we're reading
    bytesRead := make([]byte, 35)

    // Now read some data, passing in our byte slice
    n, err := file.Read(bytesRead)
    if err != nil {
        log.Fatalf("error reading from sloth.txt: %v", err)
    }

    // Take a look at the bytes we copied into the byte slice
    log.Printf("We read \"%s\" into bytesRead (%d bytes)",
        string(bytesRead), n)
}