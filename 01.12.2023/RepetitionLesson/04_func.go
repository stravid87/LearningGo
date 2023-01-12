package main

import (
	"fmt"
 	"math"
)

type circle struct {
	radius float64
}

type shape struct{
	side_a int
	side_b int
}

func (sh shape) info() int {
	return sh.side_a*sh.side_b
}

func (c circle)info() float64 {
	return math.Pi*c.radius*c.radius
}

func main() {
	sh := shape{
		side_a: 4,
		side_b: 6,
	};
	c := circle{
		radius: 2,
	}
	fmt.Println("Area of square", sh.info())
	fmt.Println("Area of circle", c.info())
}