package main

import "fmt"

type people struct {
	name                string
	last_name           string
	favourite_ice_cream []string
}

func main() {

	map_ := map[string]people{
		"Junior": people{"Ubiratan", "Junior", []string{"chocolate", "vanilla"}},
		"Costa":  people{"Jose", "Carlos", []string{"vanilla", "flakes"}},
	}

	for _, people := range map_ {
		fmt.Printf("%v:\n", people.name)

		for _, favourite_ice_cream := range people.favourite_ice_cream {
			fmt.Printf("\t%v\n", favourite_ice_cream)
		}
	}
}
