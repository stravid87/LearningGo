package polymorph

import (
	"fmt"
	"math"
)

type Shape interface {
	area()
}

type Rectangle struct {
	X1, Y1, X2, Y2 float64
}

type Circle struct {
	Xc, Yc, Radius float64
}

func (r *Rectangle) area() {
	fmt.Println("Rectangle Area : ", (r.X2-r.X1)*(r.Y2-r.Y1))
}

func (c *Circle) area() {
	fmt.Println("Circle Area: ", math.Pi*math.Pow(c.Radius, 2))
}

func GetArea(s Shape) {
	s.area()
}