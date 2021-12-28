Exercise description:

1. Create a function that returns a `int`
1. Create another function that returns a `string` and an `int`
1. Call the 2 functions
1. Show the results

> Code:
```go
package main

import "fmt"

func main() {
	x := returnInt()
	y, z := returnStringInt()

	fmt.Println(x)
	fmt.Println(y, z)
}

func returnInt() int {
	return 1
}

func returnStringInt() (string, int) {
	return "string", 1
}

```

> Output:
```console
1
string 1
```