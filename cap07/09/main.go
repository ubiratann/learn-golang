package main

import "fmt"

func main() {
	esporteFavorito := "soccer"

	switch esporteFavorito {
	case "basketball":
		fmt.Println("Your favourite sport is basketball")
	case "soccer":
		fmt.Println("Your favourite sport is soccer")
	default:
		fmt.Println("Sorry I don't know this sport :(")
	}
}
