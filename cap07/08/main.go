package main

import "fmt"

func main() {
	x := "room"

	switch {
	case x == "kitchen":
		fmt.Println("You are in the kitchen")
	case x == "room":
		fmt.Println("You are in the room")
	default:
		fmt.Println("You are not in house")
	}
}
