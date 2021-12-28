Exercise description:

1. Create a function that returns a function.
1. Assign the returned function to a variable.
1. Call the returned function. 

> Code:
```go
package main

import "fmt"

func getPower() func(x int) int {
	return power
}

func power(x int) int {
	return 2 * x
}

func main() {
	power_ := getPower()
	fmt.Println(power_(2))
}

```

> Output:
```console
4
```