package main

import (
	"fmt"
	"oop_packages/payroll"
)

func main() {
	e1 := payroll.Employee{Id: 101}
	e1.SetEmployee("Bob", "82347783", "bob@gmail.com")
	fmt.Println(e1.GetId(), ":", e1.GetName(), ":", e1.GetPhone(), ":", e1.GetEmail())
}