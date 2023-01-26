package main

import "oop_packages/polymorph"

func main() {
	r := polymorph.Rectangle{10, 10, 20, 20}
	polymorph.GetArea(&r)
	c := polymorph.Circle{5, 5, 30}
	polymorph.GetArea(&c)

}
