package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	f, err := os.Create("log.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	log.SetOutput(f) // Key point

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("error happened", err)
	}
	defer f2.Close()

	fmt.Println("check the Log TXT file in the directory")
}