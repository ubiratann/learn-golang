package main

import "fmt"

func main() {
	x := struct {
		name             string
		parents          map[string]string
		favourite_colors []string
	}{"Ubiratan Jr.", map[string]string{"Father": "Ubiratan", "Mother": "Fatima"}, []string{"red", "blue", "green"}}

	fmt.Printf("Name: %v\n", x.name)

	fmt.Printf("\tParents:\n")
	for key, value := range x.parents {
		fmt.Printf("\t\t%v: %v\n", key, value)
	}

	fmt.Println("Favourite colors:")
	for index, color := range x.favourite_colors {
		fmt.Printf("\t%v. %v\n", index, color)
	}
}
