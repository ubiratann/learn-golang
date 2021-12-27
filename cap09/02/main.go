package main

import "fmt"

func main() {
	slice := []int{}

	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)

	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Printf("%T\n", slice)
}
