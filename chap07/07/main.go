package main

import "fmt"

func main() {
	x := 20

	if x == 10 {
		fmt.Println("The 'x' number is 10")
	} else if x%2 == 0 {
		fmt.Println("The 'x' number is a even number and isn't 10")
	} else {
		fmt.Println("The 'x' number isn't 10 or a even number")
	}
}
