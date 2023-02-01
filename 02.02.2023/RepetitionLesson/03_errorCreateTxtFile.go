package main

import (
	"fmt"
	"os"
	"io"
	"strings"
)

func main() {
	var firstname, surname string
	//Creating new text.txt file
	f, err := os.Create("text.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer f.Close()

	//Fill in the firstname
	fmt.Print("Firstname: ")
	_, err = fmt.Scan(&firstname)
	if err != nil {
		fmt.Println(err)
	}

	//fill in the surname
	fmt.Print("Surname: ")
	_, err = fmt.Scan(&surname)
	if err != nil {
		fmt.Println(err)
	}	
	
	// Reader
	r := strings.NewReader(firstname)
	s := strings.NewReader(surname)
	
	//Stored to text.txt file
	io.Copy(f, r)
	io.Copy(f, s)
}