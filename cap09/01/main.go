package main

import "fmt"

func main() {
	array := [5]int{0, 0, 0, 0, 0}

	array[0] = 1
	array[1] = 2
	array[2] = 3

	for _, value := range array {
		fmt.Println(value)
	}
	fmt.Printf("%T\n", array)
}
