package main

import "fmt"

func getPower() func(x int) int {
	return power
}

func power(x int) int {
	return 2 * x
}

func main() {
	power_ := getPower()
	fmt.Println(power_(2))
}
