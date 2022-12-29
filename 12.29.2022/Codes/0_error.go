package main

import (  
    "fmt"
    "os"
)

func main() {  
    f, err := os.Open("/12.29.2022 Golang Lesson Instruction.docx")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(f.Name(), "opened successfully")
}