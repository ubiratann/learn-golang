Exercise description:

1. Create a `map` with `string` key and `[]string` value:
	- the keys could have names using the surname_name format
	- the values are favourite hobbies of the people
1. Show all this values and indexes

> Code:
```go
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

```

> Output:
```console
Ubiratan_Junior:
        0: play games
        1: listen music
Jose_Pereira:
        0: play soccer
        1: paint
```