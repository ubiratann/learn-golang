package main

import "fmt"

func main() {
	map_ := map[string][]string{}

	map_["Ubiratan_Junior"] = []string{"play games", "listen music"}
	map_["Jose_Pereira"] = []string{"play soccer", "paint"}

	for key, value := range map_ {
		fmt.Printf("%v:\n", key)
		for index, v := range value {
			fmt.Printf("\t%v: %v\n", index, v)
		}
	}
}
