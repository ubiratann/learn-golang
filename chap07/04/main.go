package main

import "fmt"

func main() {
	year := 1999
	for {
		fmt.Println(year)
		if year == 2021 {
			break
		}
		year++

	}
}
