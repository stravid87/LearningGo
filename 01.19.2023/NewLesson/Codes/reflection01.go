package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId        int
	customerId   uint
	customerName string
	customerAddr addr
}

type addr struct {
	street string
	number int
}

func introspectData(d interface{}, count int) {
	// Take any data and extract the "type" and "kind" the "type" belongs to.
	t := reflect.TypeOf(d)
	k := t.Kind()
	fmt.Println("Type: ", t) // main.order (Println uses the default format under the hood.)
	fmt.Println("Kind: ", k) // struct (note, all main.order's are of the kind, struct.)
	fmt.Println("*", count, "--------------------------------------------")

	// If the data passed in is a struct, dive into into it to extract nested structures
	if reflect.TypeOf(d).Kind().String() != "struct" {
		return
	}

	// Step 1, get the value of the data passed in
	val := reflect.ValueOf(d)

	// Step 2, get the type of the value passed in
	//valType := val.Type() // Refers to main.order

	// Step 3, run a for loop bounded by the number of fields of the struct passed in as the data
	for i := 0; i < reflect.ValueOf(d).NumField(); i++ {
		fmt.Printf("---* %d *---\n", i)
		nestedValue := val.Field(i)
		fmt.Println("Type: ", nestedValue.Type())        // int | uint | string | main.addr
		fmt.Println("Kind: ", nestedValue.Type().Kind()) // int | unint | string | struct
	}
}

func main() {
	o := order{
		ordId:        456,
		customerId:   56,
		customerName: "Ravi",
		customerAddr: addr{
			"Avenue",
			99,
		},
	}

	introspectData(o, 0)

}
