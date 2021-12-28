package main

import "fmt"

func example() func() int {
	timesHere := 0

	return func() int {
		timesHere++
		return timesHere
	}
}

func main() {
	x := example()
	y := example()
	z := x
	x()
	x()
	x()
	fmt.Println("x:", x())
	fmt.Println("y:", y())
	fmt.Println("z:", z())

}
