package main

import "fmt"

func main() {
	power := func(x int) int {
		return x * x
	}

	fmt.Println(power(10))
}
