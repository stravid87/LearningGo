package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()                     // order | employee | int
		query := fmt.Sprintf("insert into %s values(", t) // e.g., insert order values(...

		v := reflect.ValueOf(q) // v is of type reflect.Value

		for i := 0; i < v.NumField(); i++ { // for loop to iterate through all fields of the data
			switch v.Field(i).Kind() { // switch based on kind
			case reflect.Int: // if of type int
				if i == 0 { // You're at the beginning of the line above and no extra space after the trailing ( is necessary
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else { //append "comma space digit" to the growing string
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 { // append "<string>" with no added comma or space
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query) // output the final query built
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o) //insert into order values(456, 56)

	e := employee{
		name:    "Javokhir",
		id:      565,
		address: "Tashkent",
		salary:  90000,
		country: "Uzbekistan",
	}
	createQuery(e) //insert into employee values("javokhir", 565, "Taskent", 90000, "uzebekistan")

	i := 90
	createQuery(i) //unsuported type

}
