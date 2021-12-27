Exercise description:

1. Create a program that uses the switch statement, where the switch statement is a string variable with the identifier "favouriteSport".


> Code:
```go
package main

import "fmt"

func main() {
	favouriteSport := "swimming"

	switch favouriteSport {
	case "basketball":
		fmt.Println("Your favourite sport is basketball")
	case "soccer":
		fmt.Println("Your favourite sport is soccer")
	default:
		fmt.Println("Sorry I don't know this sport :(")
	}
}
```

> Output:
```console
Sorry I don't know this sport :(
```