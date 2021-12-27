package main

import "fmt"

func main() {
	slice := [][]string{
		{"Ubiratan", "Júnior", "listen trap"},
		{"Jose", "Pereira", "go to shopping"},
		{"Antonio", "Mendes", "play soccer"},
	}

	for _, value := range slice {
		fmt.Println(value[0])
		for _, value2 := range value {
			fmt.Printf("\t%v\n", value2)
		}
	}
}
