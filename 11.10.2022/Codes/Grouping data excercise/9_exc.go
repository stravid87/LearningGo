package main

import "fmt"

func main(){
	m := map[string][]string{
		`bond_james`: []string{`shaken`, `Miss`, `Moneypenny`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr` : []string{`Being evil`, `Ice cream`, `Sunsets`}, 
	}

	fmt.Println(m)

	delete(m, `no_dr`)

	m[`javokhir`] = []string{`computer Science`, `Golang`, `Vue`}

	for k, v := range m {
		fmt.Println("Index: ", k)
		for i, v2 := range v{
			fmt.Println("\t",i, v2)
		}
	}
}