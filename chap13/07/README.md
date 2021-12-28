Exercise description:

1. Assign a function to a variable
1. Call this function

> Code:
```go
package main

import "fmt"

func main() {
	power := func(x int) int {
		return x * x
	}

	fmt.Println(power(10))
}

```

> Output:
```console
100
```