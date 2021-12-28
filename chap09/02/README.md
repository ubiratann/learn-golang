Exercise description:
1. Using literal composite:
    - Create an slice of int type
    - Assign 10 values 
2. Use `range` to show the slice values
3. Use format printing to show the slice type

> Code:
```go
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Printf("%T\n", slice)
}
```

> Output:
```console
1
2
3
...
10
[]int
```