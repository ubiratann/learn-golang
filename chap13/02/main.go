package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}

	x := variadic(5, 6, 7, 8)
	y := variadic(slice...)
	z := noVariadic(slice)

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

}

func variadic(x ...int) int {
	sum := 0
	for _, value := range x {
		sum += value
	}
	return sum
}

func noVariadic(x []int) int {
	return variadic(x...)
}
