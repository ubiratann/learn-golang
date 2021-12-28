package main

import "fmt"

func main() {
	for x := 65; x < 91; x++ {

		fmt.Println("\n", x)
		for y := 0; y < 3; y++ {
			fmt.Printf("%#U ", x)
		}
		fmt.Println("")
	}
}
