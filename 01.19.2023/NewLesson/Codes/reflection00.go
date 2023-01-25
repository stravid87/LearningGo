package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	A := Record{
		"Shawinigan",
		-12.123,
		Secret{
			"Mihalis",
			"Tsoukalos",
		},
	}
	reflectedValueOf := reflect.ValueOf(A) // Returns the type? No! Returns value
	// Here, a reflect.value is cast as a string and printed out
	fmt.Printf("1) %s \n", reflectedValueOf.String()) // 1) <main.Record Value>
	// breaking this message down, main.Record refers to the value's type
	// value is this whole thing <main.Record Value>
	// the kind...?

	// Unaltered, the type of a reflect.value is a reflect.value
	fmt.Printf("2) %T \n", reflectedValueOf) // 2) reflect.Value

	// If the value is cast as a string,
	fmt.Printf("3) %T \n", reflectedValueOf.String()) // 3) String

	// Unaltered, if coerced, error results as and attempt is made to coerce all data into strings
	fmt.Printf("4) %s \n", reflectedValueOf) // 4) {Shawinigan %!s(float64=-12.123) {Mihalis Tsoukalos}}

	fmt.Printf("5) E.g., type int formatted as string: %s \n", 10) // 5) %!s(int=10)

	iType := reflectedValueOf.Type() // Here we are returning type.
	fmt.Println(" *** reflectedValueOf.Type() => ", iType)

	dotType := reflect.TypeOf(reflectedValueOf)
	fmt.Println(" *** reflect.TypeOf(reflectedValueOf) => ", dotType)

	// a refelect.rtype formatted as a string refers will print the value's underlying type
	fmt.Printf("6) i type: %s\n", iType) // 6) main.Record
	fmt.Printf("7) i type: %T\n", iType) // 7) *reflect.rtype

	fmt.Printf("8) The %d fields of %s are: \n", reflectedValueOf.NumField(), iType) //

	for i := 0; i < reflectedValueOf.NumField(); i++ {
		fmt.Println("<---------------------------------------->")

		// Get the name of each key for the struct
		fmt.Printf("\t%s ", iType.Field(i).Name) // Field1 | Field2 | Field 3

		// For each key, get the type of each associated value
		fmt.Printf("\twith type: %s ", reflectedValueOf.Field(i).Type()) // string | float64 | main.Secret

		// Formatted as %v, the value will be printed as it's default format
		fmt.Printf("\n 9) .Interface formatted as value: _%v_\n", reflectedValueOf.Field(i).Interface())

		// Now for kind. Each type can have different types?
		k := reflect.TypeOf(reflectedValueOf.Field(i).Interface() /*get underlying value of any type*/).Kind()
		fmt.Println("10) This is what .kind() returns: ", k) // string | float64 | struct

		// // if reflect.kind cast as a string
		// if k.String() == "struct" {
		// 	fmt.Println("(A)", reflectedValueOf.Field(i).Type())
		// }

		// if k == reflect.Struct {
		// 	fmt.Println("(B)", reflectedValueOf.Field(i).Type())
		// }

		// fmt.Println("(C)", reflectedValueOf.Field(i).Type())

	}

	fmt.Println("<---------------------------------------->")

	// Notice that reflect.TypeOf can be called on any data
	// Types are subsets of Kinds. So, the type of "main.Record" & "main.Secret" are both of kind "struct"

	B := Record{
		"exampleRecordOne",
		99.99,
		Secret{
			"secretFieldUsername",
			"secretFieldPassword",
		},
	}

	S := Secret{
		"usernameForIndependentSecret",
		"passwordForIndependentSecret",
	}

	// The types of var B & S will be different and specific. However, the kind will be the same
	fmt.Printf("B has type _%s_ and S has type _%s_, but both share a kind, _%s_ & _%s_",
		reflect.TypeOf(B).String(),
		reflect.TypeOf(S).String(),
		reflect.TypeOf(B).Kind().String(),
		reflect.TypeOf(S).Kind().String(),
	) // "B has type main.Record and S has type main.Secret, but both share a kind, struct & struct"

	//In the above example, it is very important to recognize that if reflect.TypeOf() is printed directly using PrintLn, the %v (i.e., the default format) is used and you will get *reflect.rtype. In general, the string representation of the underlying type will be desired.

}
