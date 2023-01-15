package main

import (
	"fmt"
	"sort"
)

type user struct {
	First string
	Last string
	Age int
}

type ByAge []user
func (a ByAge) Len() int 			{return len(a)}
func (a ByAge) Swap(i, j int) 		{ a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool  { return a[i].Age < a[j].Age }

type ByLast []user
func (l ByLast) Len() int 			{ return len(l) }
func (l ByLast) Swap(i, j int) 		{ l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Age < l[j].Age }

func main() {
	u1 := user{
		First: "James",
		Last: "Bond",
		Age: 37,
	}
	u2 := user{
		First: "Miss",
		Last: "Moneypeny",
		Age: 27,
	}
	u3 := user{
		First: "Javokhir",
		Last: "Nematov",
		Age: 20,
	}
	u4 := user{
		First: "Ravi",
		Last: "Seyed-Mahmoud",
		Age: 35,
	}

	users := []user{u1, u2, u3, u4}
	sort.Sort(ByAge(users))
	for _, u := range users{
		fmt.Println(u.First, u.Last, u.Age)
	}
	fmt.Println("-----------")
	sort.Sort(ByLast(users))
	for _, u := range users{
		fmt.Println(u.First, u.Last, u.Age)

	}

}