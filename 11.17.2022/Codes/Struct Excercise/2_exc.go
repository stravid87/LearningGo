package main 

import "fmt"

type vehicle struct{
	doors int
	color string
}

type sedan struct{
	vehicle
	wheel bool
}

type truck struct{
	vehicle
	wheel bool
}

func main() {
	t := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},

		wheel: true,
	}

	s := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		wheel: false,
	}

	fmt.Println(t)
	fmt.Println(s)
}