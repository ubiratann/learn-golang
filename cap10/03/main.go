package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWhellDrive bool
}

type sedan struct {
	vehicle
	luxuryModel bool
}

func main() {
	tk := truck{vehicle{2, "pink"}, true}
	sd := sedan{vehicle{4, "blue"}, false}

	fmt.Printf("Truck:\t%v\n", tk)
	fmt.Printf("Sedan:\t%v\n", sd)

	fmt.Printf("Truck color:\t%v\n", tk.color)
	fmt.Printf("Sedan color:\t%v\n", sd.color)

}
